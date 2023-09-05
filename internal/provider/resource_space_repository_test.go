package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSpaceProjectRepositoryResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccSpaceProjectRepositoryResourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jetbrains-space_project_repository.tests", "name", "unit-tests"),
				),
			},
		},
	})
}

func testAccSpaceProjectRepositoryResourceConfig() string {
	return `
resource "jetbrains-space_project_repository" "tests" {
  project_id          = "3dMkH62adtiH"
  name                = "unit-tests"
  description         = "Test description"
  default_branch_name = "master"
}
`
}
