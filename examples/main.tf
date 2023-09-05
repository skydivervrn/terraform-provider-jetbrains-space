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

resource "jetbrains-space_project" "tests" {
  name        = "Tests-ng"
  description = "Test description"
  private     = false
}

resource "jetbrains-space_project_repository" "tests" {
  project_id          = jetbrains-space_project.tests.id
  name                = "unit-tests"
  description         = "Test description"
  default_branch_name = "master"
}

output "data_test_output" {
  value = jetbrains-space_project.tests.id
}

output "data_test_output2" {
  value = jetbrains-space_project_repository.tests.default_branch_name
}