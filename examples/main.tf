terraform {
  required_providers {
    jetbrains-space = {
      #       version = "0.0.1"
      source = "registry.terraform.io/skydivervrn/jetbrains-space"
    }
  }
}

provider "jetbrains-space" {}

data "jetbrains-space_project" "this" {
  id = "19lHBY1GRU5x"
}
#data "jetbrains-space_example" "this" {}

# resource "jetbrains-space_project" "tests" {
#   name = "Tests-ng"
# }
#
# resource "jetbrains-space_git_repository" "tests" {
#   project_id          = jetbrains-space_project.tests.id
#   name                = "test"
#   default_branch_name = "master"
# }
#
# output "data_test_output" {
#   value = data.jetbrains-space_project_data_source.this.id
# }
#
# output "data_test_output2" {
#   value = jetbrains-space_project.tests.id
# }