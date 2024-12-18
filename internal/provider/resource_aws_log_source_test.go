package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAWSLogSourceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unknown log forwarder ID fails
			{
				Config: `
resource "ghost_aws_log_source" "test" {
    log_forwarder_id = "503af559-6195-4076-b1ca-15d72c7d22fe"
    s3_bucket_name = "ghost-example-bucket"
    role_arn = "arn:aws:iam::012345678901:role/ghost-f13217e5-6a54-44e0-ba67-9c24e7b0e542-ingest"
    sqs_arn = "arn:aws:sqs:us-west-2:012345678901:ghost-f13217e5-6a54-44e0-ba67-9c24e7b0e542-ingest"
    account_id = "012345678901"
    region = "us-west-2"
}
`,
				ExpectError: regexp.MustCompile(".*Log forwarder not found.*"),
			},
		},
	})
}
