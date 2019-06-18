---
layout: "aci"
page_title: "ACI: aci_vrf"
sidebar_current: "docs-aci-resource-vrf"
description: |-
  Manages ACI VRF
---

# aci_vrf #
Manages ACI VRF

## Example Usage ##

```hcl
resource "aci_vrf" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  bd_enforced_enable  = "example"
  ip_data_plane_learning  = "example"
  knw_mcast_act  = "example"
  name_alias  = "example"
  pc_enf_dir  = "example"
  pc_enf_pref  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object vrf.
* `annotation` - (Optional) annotation for object vrf.
* `bd_enforced_enable` - (Optional) bd_enforced_enable for object vrf.
* `ip_data_plane_learning` - (Optional) ip_data_plane_learning for object vrf.
* `knw_mcast_act` - (Optional) specifies if known multicast traffic is forwarded
* `name_alias` - (Optional) name_alias for object vrf.
* `pc_enf_dir` - (Optional) pc_enf_dir for object vrf.
* `pc_enf_pref` - (Optional) preferred policy control

* `relation_fv_rs_ospf_ctx_pol` - (Optional) Relation to class ospfCtxPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_vrf_validation_pol` - (Optional) Relation to class l3extVrfValidationPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_ctx_mcast_to` - (Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_ctx_to_eigrp_ctx_af_pol` - (Optional) Relation to class eigrpCtxAfPol. Cardinality - N_TO_M. Type - Set of Map.
                
* `relation_fv_rs_ctx_to_ospf_ctx_pol` - (Optional) Relation to class ospfCtxPol. Cardinality - N_TO_M. Type - Set of Map.
                
* `relation_fv_rs_ctx_to_ep_ret` - (Optional) Relation to class fvEpRetPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bgp_ctx_pol` - (Optional) Relation to class bgpCtxPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_ctx_mon_pol` - (Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_ctx_to_ext_route_tag_pol` - (Optional) Relation to class l3extRouteTagPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_ctx_to_bgp_ctx_af_pol` - (Optional) Relation to class bgpCtxAfPol. Cardinality - N_TO_M. Type - Set of Map.
                


