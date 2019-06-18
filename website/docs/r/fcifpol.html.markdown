---
layout: "aci"
page_title: "ACI: aci_interface_fc_policy"
sidebar_current: "docs-aci-resource-interface_fc_policy"
description: |-
  Manages ACI Interface FC Policy
---

# aci_interface_fc_policy #
Manages ACI Interface FC Policy

## Example Usage ##

```hcl
resource "aci_interface_fc_policy" "example" {


    name  = "example"

  annotation  = "example"
  automaxspeed  = "example"
  fill_pattern  = "example"
  name_alias  = "example"
  port_mode  = "example"
  rx_bb_credit  = "example"
  speed  = "example"
  trunk_mode  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object interface_fc_policy.
* `annotation` - (Optional) annotation for object interface_fc_policy.
* `automaxspeed` - (Optional) automaxspeed for object interface_fc_policy.
* `fill_pattern` - (Optional) fill_pattern for object interface_fc_policy.
* `name_alias` - (Optional) name_alias for object interface_fc_policy.
* `port_mode` - (Optional) 
* `rx_bb_credit` - (Optional) rx_bb_credit for object interface_fc_policy.
* `speed` - (Optional) cpu or port speed
* `trunk_mode` - (Optional) trunk_mode for object interface_fc_policy.



