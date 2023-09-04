terraform {
  required_providers {
    jetbrains-space = {
      version = "0.0.12"
      source  = "hashicorp.com/skydivervrn/jetbrains-space"
    }
  }
}

provider "jetbrains-space" {}

data "jetbrains-space_project_data_source" "this" {
  id = jetbrains-space_project.tests.id
}

resource "jetbrains-space_project" "tests" {
  name = "Tests-ng"
}

resource "jetbrains-space_git_repository" "tests" {
  project_id          = jetbrains-space_project.tests.id
  name                = "test"
  default_branch_name = "master"
}

output "data_test_output" {
  value = data.jetbrains-space_project_data_source.this.id
}

output "data_test_output2" {
  value = jetbrains-space_project.tests.id
}