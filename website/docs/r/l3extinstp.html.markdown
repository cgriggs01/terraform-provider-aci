---
layout: "aci"
page_title: "ACI: aci_external_network_instance_profile"
sidebar_current: "docs-aci-resource-external_network_instance_profile"
description: |-
  Manages ACI External Network Instance Profile
---

# aci_external_network_instance_profile #
Manages ACI External Network Instance Profile

## Example Usage ##

```hcl
resource "aci_external_network_instance_profile" "example" {

  l3_outside_dn  = "${aci_l3_outside.example.id}"

    name  = "example"

  annotation  = "example"
  exception_tag  = "example"
  flood_on_encap  = "example"
  match_t  = "example"
  name_alias  = "example"
  pref_gr_memb  = "example"
  prio  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `l3_outside_dn` - (Required) Distinguished name of parent L3Outside object.
* `name` - (Required) name of Object external_network_instance_profile.
* `annotation` - (Optional) annotation for object external_network_instance_profile.
* `exception_tag` - (Optional) exception_tag for object external_network_instance_profile.
* `flood_on_encap` - (Optional) flood_on_encap for object external_network_instance_profile.
* `match_t` - (Optional) match criteria
* `name_alias` - (Optional) name_alias for object external_network_instance_profile.
* `pref_gr_memb` - (Optional) pref_gr_memb for object external_network_instance_profile.
* `prio` - (Optional) qos priority class id
* `target_dscp` - (Optional) target dscp

* `relation_fv_rs_sec_inherited` - (Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prov` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_l3ext_rs_l3_inst_p_to_dom_p` - (Optional) Relation to class extnwDomP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_l3ext_rs_inst_p_to_nat_mapping_e_pg` - (Optional) Relation to class fvAEPg. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_cons_if` - (Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cust_qos_pol` - (Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_l3ext_rs_inst_p_to_profile` - (Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.
                
* `relation_fv_rs_cons` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prot_by` - (Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_intra_epg` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                


