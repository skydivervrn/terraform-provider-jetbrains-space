package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	testNameSpaceRepositoryResource = "go-test"
)

func TestAccRepository(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRepository(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						fmt.Sprintf("%s.test", spaceRepositoryResourceName), "name", regexp.MustCompile("")),
				),
			},
		},
	})
}

func testAccRepository() string {
	return fmt.Sprintf(`
resource %s "test" {
  project_id = "%s"
  name = "%s"
}
`,
		spaceRepositoryResourceName,
		spaceProjectDataSourceTestResourceId,
		testNameSpaceRepositoryResource,
	)
}
