---
layout: "aci"
page_title: "ACI: aci_application_epg"
sidebar_current: "docs-aci-resource-application_epg"
description: |-
  Manages ACI Application EPG
---

# aci_application_epg #
Manages ACI Application EPG

## Example Usage ##

```hcl
resource "aci_application_epg" "example" {

  application_profile_dn  = "${aci_application_profile.example.id}"

    name  = "example"

  annotation  = "example"
  exception_tag  = "example"
  flood_on_encap  = "example"
  fwd_ctrl  = "example"
  has_mcast_source  = "example"
  is_attr_based_e_pg  = "example"
  match_t  = "example"
  name_alias  = "example"
  pc_enf_pref  = "example"
  pref_gr_memb  = "example"
  prio  = "example"
  shutdown  = "example"
}
```
## Argument Reference ##
* `application_profile_dn` - (Required) Distinguished name of parent ApplicationProfile object.
* `name` - (Required) name of Object application_epg.
* `annotation` - (Optional) annotation for object application_epg.
* `exception_tag` - (Optional) exception_tag for object application_epg.
* `flood_on_encap` - (Optional) flood_on_encap for object application_epg.
* `fwd_ctrl` - (Optional) fwd_ctrl for object application_epg.
* `has_mcast_source` - (Optional) has_mcast_source for object application_epg.
* `is_attr_based_e_pg` - (Optional) is_attr_based_e_pg for object application_epg.
* `match_t` - (Optional) match criteria
* `name_alias` - (Optional) name_alias for object application_epg.
* `pc_enf_pref` - (Optional) enforcement preference
* `pref_gr_memb` - (Optional) pref_gr_memb for object application_epg.
* `prio` - (Optional) qos priority class id
* `shutdown` - (Optional) shutdown for object application_epg.

* `relation_fv_rs_bd` - (Optional) Relation to class fvBD. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_cust_qos_pol` - (Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_dom_att` - (Optional) Relation to class infraDomP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_fc_path_att` - (Optional) Relation to class fabricPathEp. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prov` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_graph_def` - (Optional) Relation to class vzGraphCont. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cons_if` - (Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_sec_inherited` - (Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_node_att` - (Optional) Relation to class fabricNode. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_dpp_pol` - (Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_cons` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prov_def` - (Optional) Relation to class vzCtrctEPgCont. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_trust_ctrl` - (Optional) Relation to class fhsTrustCtrlPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_path_att` - (Optional) Relation to class fabricPathEp. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prot_by` - (Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_ae_pg_mon_pol` - (Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_intra_epg` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                


