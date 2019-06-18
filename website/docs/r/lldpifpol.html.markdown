---
layout: "aci"
page_title: "ACI: aci_lldp_interface_policy"
sidebar_current: "docs-aci-resource-lldp_interface_policy"
description: |-
  Manages ACI LLDP Interface Policy
---

# aci_lldp_interface_policy #
Manages ACI LLDP Interface Policy

## Example Usage ##

```hcl
resource "aci_lldp_interface_policy" "example" {


    name  = "example"

  admin_rx_st  = "example"
  admin_tx_st  = "example"
  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object lldp_interface_policy.
* `admin_rx_st` - (Optional) admin receive state
* `admin_tx_st` - (Optional) admin transmit state
* `annotation` - (Optional) annotation for object lldp_interface_policy.
* `name_alias` - (Optional) name_alias for object lldp_interface_policy.



