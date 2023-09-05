package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSpaceProjectResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccSpaceProjectResourceConfig(),
				Check:  resource.ComposeAggregateTestCheckFunc(
				//resource.TestCheckResourceAttr("jetbrains-space_project.test", "name", "Unit-tests"),
				),
			},
		},
	})
}

func testAccSpaceProjectResourceConfig() string {
	return fmt.Sprintf(`
resource "jetbrains-space_project" "test" {
  name        = "Unit-tests-%s"
  description = "Unit-tests"
  private     = false
}
`, testSuffix)
}
