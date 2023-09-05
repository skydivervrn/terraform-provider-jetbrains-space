package provider

import (
	"fmt"
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
				Check:  resource.ComposeAggregateTestCheckFunc(
				//resource.TestCheckResourceAttr("jetbrains-space_project_repository.tests", "name", "unit-tests"),
				),
			},
		},
	})
}

func testAccSpaceProjectRepositoryResourceConfig() string {
	return fmt.Sprintf(`
resource "jetbrains-space_project_repository" "tests" {
  project_id          = "%s"
  name                = "unit-tests-%s"
  description         = "Test description"
  default_branch_name = "master"
}
`, testProjectId, testSuffix)
}
