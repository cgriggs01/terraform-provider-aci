---
layout: "aci"
page_title: "ACI: aci_ospf_interface_policy"
sidebar_current: "docs-aci-resource-ospf_interface_policy"
description: |-
  Manages ACI OSPF Interface Policy
---

# aci_ospf_interface_policy #
Manages ACI OSPF Interface Policy

## Example Usage ##

```hcl
resource "aci_ospf_interface_policy" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  cost  = "example"
  ctrl  = "example"
  dead_intvl  = "example"
  hello_intvl  = "example"
  name_alias  = "example"
  nw_t  = "example"
  pfx_suppress  = "example"
  prio  = "example"
  rexmit_intvl  = "example"
  xmit_delay  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object ospf_interface_policy.
* `annotation` - (Optional) annotation for object ospf_interface_policy.
* `cost` - (Optional) ospf interface cost
* `ctrl` - (Optional) interface policy controls
* `dead_intvl` - (Optional) interval between hello packets
* `hello_intvl` - (Optional) interface policy hello interval
* `name_alias` - (Optional) name_alias for object ospf_interface_policy.
* `nw_t` - (Optional) ospf interface network type
* `pfx_suppress` - (Optional) pfx_suppress for object ospf_interface_policy.
* `prio` - (Optional) ospf interface priority
* `rexmit_intvl` - (Optional) interval between lsa retransmissions
* `xmit_delay` - (Optional) delay time to send an lsa update packet



