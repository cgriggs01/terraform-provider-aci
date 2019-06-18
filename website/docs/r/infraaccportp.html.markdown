---
layout: "aci"
page_title: "ACI: aci_leaf_interface_profile"
sidebar_current: "docs-aci-resource-leaf_interface_profile"
description: |-
  Manages ACI Leaf Interface Profile
---

# aci_leaf_interface_profile #
Manages ACI Leaf Interface Profile

## Example Usage ##

```hcl
resource "aci_leaf_interface_profile" "example" {


    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object leaf_interface_profile.
* `annotation` - (Optional) annotation for object leaf_interface_profile.
* `name_alias` - (Optional) name_alias for object leaf_interface_profile.



