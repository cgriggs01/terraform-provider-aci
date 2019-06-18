---
layout: "aci"
page_title: "ACI: aci_vmm_domain"
sidebar_current: "docs-aci-resource-vmm_domain"
description: |-
  Manages ACI VMM Domain
---

# aci_vmm_domain #
Manages ACI VMM Domain

## Example Usage ##

```hcl
resource "aci_vmm_domain" "example" {

  provider_profile_dn  = "${aci_provider_profile.example.id}"

    name  = "example"

  access_mode  = "example"
  annotation  = "example"
  arp_learning  = "example"
  ave_time_out  = "example"
  config_infra_pg  = "example"
  ctrl_knob  = "example"
  delimiter  = "example"
  enable_ave  = "example"
  enable_tag  = "example"
  encap_mode  = "example"
  enf_pref  = "example"
  ep_inventory_type  = "example"
  ep_ret_time  = "example"
  hv_avail_monitor  = "example"
  mcast_addr  = "example"
  mode  = "example"
  name_alias  = "example"
  pref_encap_mode  = "example"
}
```
## Argument Reference ##
* `provider_profile_dn` - (Required) Distinguished name of parent ProviderProfile object.
* `name` - (Required) name of Object vmm_domain.
* `access_mode` - (Optional) access_mode for object vmm_domain.
* `annotation` - (Optional) annotation for object vmm_domain.
* `arp_learning` - (Optional) arp_learning for object vmm_domain.
* `ave_time_out` - (Optional) ave_time_out for object vmm_domain.
* `config_infra_pg` - (Optional) config_infra_pg for object vmm_domain.
* `ctrl_knob` - (Optional) ctrl_knob for object vmm_domain.
* `delimiter` - (Optional) delimiter for object vmm_domain.
* `enable_ave` - (Optional) enable_ave for object vmm_domain.
* `enable_tag` - (Optional) enable_tag for object vmm_domain.
* `encap_mode` - (Optional) encap_mode for object vmm_domain.
* `enf_pref` - (Optional) switching enforcement preference
* `ep_inventory_type` - (Optional) ep_inventory_type for object vmm_domain.
* `ep_ret_time` - (Optional) ep_ret_time for object vmm_domain.
* `hv_avail_monitor` - (Optional) hv_avail_monitor for object vmm_domain.
* `mcast_addr` - (Optional) multicast address
* `mode` - (Optional) switch used for the domain profile
* `name_alias` - (Optional) name_alias for object vmm_domain.
* `pref_encap_mode` - (Optional) pref_encap_mode for object vmm_domain.

* `relation_vmm_rs_pref_enhanced_lag_pol` - (Optional) Relation to class lacpEnhancedLagPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_vlan_ns` - (Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_dom_mcast_addr_ns` - (Optional) Relation to class fvnsMcastAddrInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_cdp_if_pol` - (Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_lacp_lag_pol` - (Optional) Relation to class lacpLagPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_vlan_ns_def` - (Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_vip_addr_ns` - (Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_lldp_if_pol` - (Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_stp_if_pol` - (Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_dom_vxlan_ns_def` - (Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_fw_pol` - (Optional) Relation to class nwsFwPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_l2_inst_pol` - (Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String.
                


