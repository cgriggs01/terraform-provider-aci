---
layout: "aci"
page_title: "ACI: aci_l2_interface_policy"
sidebar_current: "docs-aci-resource-l2_interface_policy"
description: |-
  Manages ACI L2 Interface Policy
---

# aci_l2_interface_policy #
Manages ACI L2 Interface Policy

## Example Usage ##

```hcl
resource "aci_l2_interface_policy" "example" {


    name  = "example"

  annotation  = "example"
  name_alias  = "example"
  qinq  = "example"
  vepa  = "example"
  vlan_scope  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object l2_interface_policy.
* `annotation` - (Optional) annotation for object l2_interface_policy.
* `name_alias` - (Optional) name_alias for object l2_interface_policy.
* `qinq` - (Optional) qinq for object l2_interface_policy.
* `vepa` - (Optional) vepa for object l2_interface_policy.
* `vlan_scope` - (Optional) l2 interface vlan scope



