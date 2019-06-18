---
layout: "aci"
page_title: "ACI: aci_miscabling_protocol_interface_policy"
sidebar_current: "docs-aci-resource-miscabling_protocol_interface_policy"
description: |-
  Manages ACI Mis-cabling Protocol Interface Policy
---

# aci_miscabling_protocol_interface_policy #
Manages ACI Mis-cabling Protocol Interface Policy

## Example Usage ##

```hcl
resource "aci_miscabling_protocol_interface_policy" "example" {


    name  = "example"

  admin_st  = "example"
  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object miscabling_protocol_interface_policy.
* `admin_st` - (Optional) administrative state of the object or policy
* `annotation` - (Optional) annotation for object miscabling_protocol_interface_policy.
* `name_alias` - (Optional) name_alias for object miscabling_protocol_interface_policy.



