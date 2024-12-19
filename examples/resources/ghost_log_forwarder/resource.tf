resource "ghost_log_forwarder" "forwarder" {
  name = "dev-alb-forwarder"
}

# Create a IAM policy that will allow the Ghost platform
# to assume a role in order to copy logs for processing.
data "aws_iam_policy_document" "assume_role_policy" {
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
