---
layout: "aci"
page_title: "ACI: aci_subnet"
sidebar_current: "docs-aci-resource-subnet"
description: |-
  Manages ACI Subnet
---

# aci_subnet #
Manages ACI Subnet

## Example Usage ##

```hcl
resource "aci_subnet" "example" {

  bridge_domain_dn  = "${aci_bridge_domain.example.id}"

    ip  = "example"

  annotation  = "example"
  ctrl  = "example"
  ip  = "example"
  name_alias  = "example"
  preferred  = "example"
  scope  = "example"
  virtual  = "example"
}
```
## Argument Reference ##
* `bridge_domain_dn` - (Required) Distinguished name of parent BridgeDomain object.
* `ip` - (Required) ip of Object subnet.
* `annotation` - (Optional) annotation for object subnet.
* `ctrl` - (Optional) subnet control state
* `ip` - (Optional) default gateway IP address and mask
* `name_alias` - (Optional) name_alias for object subnet.
* `preferred` - (Optional) subnet preferred status
* `scope` - (Optional) subnet visibility
* `virtual` - (Optional) virtual for object subnet.

* `relation_fv_rs_bd_subnet_to_out` - (Optional) Relation to class l3extOut. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_nd_pfx_pol` - (Optional) Relation to class ndPfxPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_subnet_to_profile` - (Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String.
                


