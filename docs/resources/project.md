---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jetbrains-space_project Resource - terraform-provider-jetbrains-space"
subcategory: ""
description: |-
  
---

# jetbrains-space_project (Resource)



## Example Usage

```terraform
resource "jetbrains-space_project" "tests" {
  name        = "Tests-ng"
  description = "Test description"
  private     = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of a project

### Optional

- `description` (String) Project description
- `private` (Boolean) Defines if project private or public

### Read-Only

- `archived` (Boolean) TBD
- `created_at_iso` (String) TBD
- `created_at_timestamp` (Number) TBD
- `icon` (String) TBD
- `id` (String) The uniq ID of existed Space project to get data from
- `key` (String) Magic field. Project name written in uppercase
- `latest_repository_activity` (String) Latest activity in project repos
