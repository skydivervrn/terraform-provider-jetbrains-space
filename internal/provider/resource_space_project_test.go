package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	spaceProjectResourceTestName = "space_project_resource_test"
)

func TestAccResourceSpace(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSpaceProject(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						fmt.Sprintf("%s.test", spaceProjectResourceName), "id", regexp.MustCompile("")),
				),
			},
		},
	})
}

func testAccSpaceProject() string {
	return fmt.Sprintf(`
resource %s "test" {
  name = "%s"
}
`,
		spaceProjectResourceName,
		spaceProjectResourceTestName,
	)
}
