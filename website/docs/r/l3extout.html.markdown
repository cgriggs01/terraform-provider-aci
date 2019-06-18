---
layout: "aci"
page_title: "ACI: aci_l3_outside"
sidebar_current: "docs-aci-resource-l3_outside"
description: |-
  Manages ACI L3 Outside
---

# aci_l3_outside #
Manages ACI L3 Outside

## Example Usage ##

```hcl
resource "aci_l3_outside" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  enforce_rtctrl  = "example"
  name_alias  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object l3_outside.
* `annotation` - (Optional) annotation for object l3_outside.
* `enforce_rtctrl` - (Optional) enforce route control type
* `name_alias` - (Optional) name_alias for object l3_outside.
* `target_dscp` - (Optional) target dscp

* `relation_l3ext_rs_dampening_pol` - (Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.
                
* `relation_l3ext_rs_ectx` - (Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.
                
* `relation_l3ext_rs_out_to_bd_public_subnet_holder` - (Optional) Relation to class fvBDPublicSubnetHolder. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_l3ext_rs_interleak_pol` - (Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String.
                
* `relation_l3ext_rs_l3_dom_att` - (Optional) Relation to class extnwDomP. Cardinality - N_TO_ONE. Type - String.
                


