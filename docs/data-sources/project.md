---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jetbrains-space_project Data Source - terraform-provider-jetbrains-space"
subcategory: ""
description: |-
  
---

# jetbrains-space_project (Data Source)



## Example Usage

```terraform
data "jetbrains-space" "this" {
  id = "YOUR_PROJECT_ID"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The uniq ID of existed Space project to get data from

### Read-Only

- `archived` (Boolean) TBD
- `created_at_iso` (String) TBD
- `created_at_timestamp` (Number) TBD
- `description` (String) Project description
- `icon` (String) TBD
- `key` (String) Magic field. Project name written in uppercase
- `latest_repository_activity` (String) Latest activity in project repos
- `name` (String) Name of a project
- `private` (Boolean) Defines if project private or public
