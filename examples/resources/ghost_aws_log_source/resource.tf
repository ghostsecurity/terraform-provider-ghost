resource "ghost_log_forwarder" "forwarder" {
  name = "dev-alb-forwarder"
}

# Connect the log forwarder to an S3 bucket with SQS notifications
# to enable the Ghost platform to read new log files as they are created
resource "ghost_aws_log_source" "source" {
  log_forwarder_id = ghost_log_forwarder.forwarder.id
  s3_bucket_name   = data.aws_s3_bucket.source.id
  role_arn         = aws_iam_role.ghost.arn
  sqs_arn          = aws_sqs_queue.notifications.arn
  account_id       = data.aws_caller_identity.current.account_id
  region           = data.aws_region.current.name

  depends_on = [
    aws_sqs_queue_policy.source_notifications,
    aws_iam_role_policy_attachment.platform_access,
  ]
}

# AWS resources to provide access to logs from the Ghost platform

# S3 bucket that is the source of log files
data "aws_s3_bucket" "source" {
  bucket = "bucket-with-logs"
}

data "aws_caller_identity" "current" {}

# SQS queue for new object notifications in source bucket
resource "aws_sqs_queue" "notifications" {
  sqs_managed_sse_enabled = true
  name                    = "ghost-notifications"

  tags = local.tags
}

data "aws_iam_policy_document" "s3_object_notification" {
  statement {
    actions = ["sqs:SendMessage"]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceAccount"
      values = [
        data.aws_caller_identity.current.account_id,
      ]
    }
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["s3.amazonaws.com"]
    }
    resources = [
      aws_sqs_queue.notifications.arn,
    ]
  }
}

# Attach policy to queue
resource "aws_sqs_queue_policy" "source_notifications" {
  policy    = data.aws_iam_policy_document.s3_object_notification.json
  queue_url = aws_sqs_queue.notifications.id
}

# Send notifications to queue for new files in bucket
resource "aws_s3_bucket_notification" "source" {
  bucket = data.aws_s3_bucket.source.id

  queue {
    queue_arn = aws_sqs_queue.notifications.arn
    events    = ["s3:ObjectCreated:*"]
  }
  depends_on = [
    aws_sqs_queue_policy.source_notifications,
  ]
}

# Create an IAM policy that will allow the Ghost platform
# to assume a role in order to copy logs for processing.
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    principals {
      type        = "Federated"
      identifiers = ["accounts.google.com"]
    }
    condition {
      test     = "Null"
      variable = "accounts.google.com:sub"
      values   = [false]
    }
    condition {
      test     = "StringEquals"
      variable = "accounts.google.com:sub"
      values   = [ghost_log_forwarder.forwarder.subject_id]
    }
    effect = "Allow"
  }
}

# Create an IAM policy that will allow the role assumed by
# the Ghost platform to manage files in the source bucket and
# receive event notifications for new files.
data "aws_iam_policy_document" "log_ingest" {
  statement {
    actions = [
      "s3:GetBucketLocation",
      "s3:GetObject",
      "s3:GetObjectAttributes",
      "s3:ListBucket"
    ]
    resources = [
      data.aws_s3_bucket.source.arn,
      "${data.aws_s3_bucket.source.arn}/*"
    ]
    effect = "Allow"
  }

  statement {
    actions = [
      "sqs:DeleteMessage",
      "sqs:ChangeMessageVisibility",
      "sqs:ReceiveMessage"
    ]
    resources = [
      aws_sqs_queue.notifications.arn
    ]
    effect = "Allow"
  }

  statement {
    actions = [
      "s3:DeleteObject"
    ]
    resources = [
      "${data.aws_s3_bucket.source.arn}/*"
    ]
    effect = "Allow"
  }
}

# Role that will be assumed by the Ghost platform
resource "aws_iam_role" "ghost" {
  name               = "ghost-ingest"
  description        = "Allows read/write access to the bucket"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  force_detach_policies = true
}

resource "aws_iam_policy" "log_ingest" {
  name        = "ghost-manage-logs"
  description = "IAM policy to allow reading logs from the bucket and receiving notifications from the SQS queue"
  policy      = data.aws_iam_policy_document.log_ingest.json
}

resource "aws_iam_role_policy_attachment" "platform_access" {
  policy_arn = aws_iam_policy.log_ingest.arn
  role       = aws_iam_role.log_ingest.name
}
