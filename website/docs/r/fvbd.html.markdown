---
layout: "aci"
page_title: "ACI: aci_bridge_domain"
sidebar_current: "docs-aci-resource-bridge_domain"
description: |-
  Manages ACI Bridge Domain
---

# aci_bridge_domain #
Manages ACI Bridge Domain

## Example Usage ##

```hcl
resource "aci_bridge_domain" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  optimize_wan_bandwidth  = "example"
  annotation  = "example"
  arp_flood  = "example"
  ep_clear  = "example"
  ep_move_detect_mode  = "example"
  host_based_routing  = "example"
  intersite_bum_traffic_allow  = "example"
  intersite_l2_stretch  = "example"
  ip_learning  = "example"
  ipv6_mcast_allow  = "example"
  limit_ip_learn_to_subnets  = "example"
  ll_addr  = "example"
  mac  = "example"
  mcast_allow  = "example"
  multi_dst_pkt_act  = "example"
  name_alias  = "example"
  type  = "example"
  unicast_route  = "example"
  unk_mac_ucast_act  = "example"
  unk_mcast_act  = "example"
  v6unk_mcast_act  = "example"
  vmac  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object bridge_domain.
* `optimize_wan_bandwidth` - (Optional) optimize_wan_bandwidth for object bridge_domain.
* `annotation` - (Optional) annotation for object bridge_domain.
* `arp_flood` - (Optional) arp flood enable
* `ep_clear` - (Optional) ep_clear for object bridge_domain.
* `ep_move_detect_mode` - (Optional) ep move detection garp based mode
* `host_based_routing` - (Optional) enables advertising host routes out of l3outs of this BD
* `intersite_bum_traffic_allow` - (Optional) 
* `intersite_l2_stretch` - (Optional) 
* `ip_learning` - (Optional) Endpoint Dataplane Learning
* `ipv6_mcast_allow` - (Optional) ipv6_mcast_allow for object bridge_domain.
* `limit_ip_learn_to_subnets` - (Optional) limits ip learning to bd subnets only
* `ll_addr` - (Optional) override of system generated ipv6 link-local address
* `mac` - (Optional) mac address
* `mcast_allow` - (Optional) mcast_allow for object bridge_domain.
* `multi_dst_pkt_act` - (Optional) forwarding method for multi destinations
* `name_alias` - (Optional) name_alias for object bridge_domain.
* `type` - (Optional) component type
* `unicast_route` - (Optional) Unicast routing
* `unk_mac_ucast_act` - (Optional) forwarding method for l2 destinations
* `unk_mcast_act` - (Optional) parameter used by node to forward data
* `v6unk_mcast_act` - (Optional) v6unk_mcast_act for object bridge_domain.
* `vmac` - (Optional) vmac for object bridge_domain.

* `relation_fv_rs_bd_to_profile` - (Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_mldsn` - (Optional) Relation to class mldSnoopPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_abd_pol_mon_pol` - (Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_to_nd_p` - (Optional) Relation to class ndIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_flood_to` - (Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_bd_to_fhs` - (Optional) Relation to class fhsBDPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_to_relay_p` - (Optional) Relation to class dhcpRelayP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_ctx` - (Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_to_netflow_monitor_pol` - (Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.
                
* `relation_fv_rs_igmpsn` - (Optional) Relation to class igmpSnoopPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_to_ep_ret` - (Optional) Relation to class fvEpRetPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_bd_to_out` - (Optional) Relation to class l3extOut. Cardinality - N_TO_M. Type - Set of String.
                


