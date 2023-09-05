// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var (
	testId = os.Getenv("TERRAFORM_PROVIDER_JETBRAINS_SPACE_TESTS_PROJECT_ID")
)

func TestAccSpaceProjectDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccSpaceProjectDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.jetbrains-space_project.test", "id", testId),
				),
			},
		},
	})
}

func testAccSpaceProjectDataSourceConfig() string {
	return fmt.Sprintf(`
data "jetbrains-space_project" "test" {
    id = "%s"
}
`,
		testId,
	)
}
