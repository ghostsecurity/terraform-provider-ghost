package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLogForwarderResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: `
resource "ghost_log_forwarder" "test" {
    name = "test-forwarder"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name is set correctly
					resource.TestCheckResourceAttr("ghost_log_forwarder.test", "name", "test-forwarder"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("ghost_log_forwarder.test", "id"),
					resource.TestCheckResourceAttrSet("ghost_log_forwarder.test", "subject_id"),
				),
			},
			// Update and Read
			{
				Config: `
resource "ghost_log_forwarder" "test" {
    name = "test-forwarder-renamed"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify first order item updated
					resource.TestCheckResourceAttr("ghost_log_forwarder.test", "name", "test-forwarder-renamed"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("ghost_log_forwarder.test", "id"),
					resource.TestCheckResourceAttrSet("ghost_log_forwarder.test", "subject_id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
