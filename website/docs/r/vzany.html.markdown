---
layout: "aci"
page_title: "ACI: aci_any"
sidebar_current: "docs-aci-resource-any"
description: |-
  Manages ACI Any
---

# aci_any #
Manages ACI Any

## Example Usage ##

```hcl
resource "aci_any" "example" {

  vrf_dn  = "${aci_vrf.example.id}"
  annotation  = "example"
  match_t  = "example"
  name_alias  = "example"
  pref_gr_memb  = "example"
}
```
## Argument Reference ##
* `vrf_dn` - (Required) Distinguished name of parent VRF object.
* `annotation` - (Optional) annotation for object any.
* `match_t` - (Optional) label match criteria
* `name_alias` - (Optional) name_alias for object any.
* `pref_gr_memb` - (Optional) pref_gr_memb for object any.

* `relation_vz_rs_any_to_cons` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_vz_rs_any_to_cons_if` - (Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_vz_rs_any_to_prov` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                


