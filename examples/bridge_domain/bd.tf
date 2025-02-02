resource "aci_tenant" "tenant_for_bridge_domain" {
  name        = "tenant_for_bd"
  description = "This tenant is created by terraform ACI provider"
}

resource "aci_bridge_domain" "demobd" {
  tenant_dn                      = "${aci_tenant.tenant_for_bridge_domain.id}"
  name                           = "test_tf_bd"
  description                    = "This bridge domain is created by terraform ACI provider"
  mac                            = "00:22:BD:F8:19:FF"
  optimize_wan_bandwidth         = "no"
  arp_flood                      = "no"
  ep_clear                       = "no"
  ep_move_detect_mode            = "garp"
  intersite_bum_traffic_allow    = "yes"
  intersite_l2_stretch           = "yes"
  ip_learning                    = "yes"
  limit_ip_learn_to_subnets      = "yes"
  mcast_allow                    = "yes"
  multi_dst_pkt_act              = "bd-flood"
  type                           = "regular"
  unicast_route                  = "no"
  unk_mac_ucast_act              = "flood"
  unk_mcast_act                  = "flood"
  vmac                           = "not-applicable"
  relation_fv_rs_bd_to_profile   = "${aci_rest.rest_rt_ctrl_profile.id}" # Relation to rtctrlProfile class. Cardinality - N_TO_ONE
  relation_fv_rs_bd_to_relay_p   = "${aci_rest.rest_dhcp_RelayP.id}"     # Relation to dhcpRelayP class. Cardinality - N_TO_ONE
  relation_fv_rs_abd_pol_mon_pol = "${aci_rest.rest_mon_epg_pol.id}"     # Relation to monEPGPol class. Cardinality - N_TO_ONE
  relation_fv_rs_bd_flood_to     = ["${aci_filter.bd_flood_filter.id}"]  # Relation to vzFilter class. Cardinality - N_TO_M
  relation_fv_rs_bd_to_fhs       = "${aci_rest.rest_fhs_bd_pol.id}"      # Relation to fhsBDPol class. Cardinality - N_TO_ONE.

  relation_fv_rs_bd_to_netflow_monitor_pol {
    tn_netflow_monitor_pol_name = "${aci_rest.rest_net_flow_pol.id}"
    flt_type                    = "ipv4"
  } # Relation to netflowMonitorPol class. Cardinality - N_TO_M

  relation_fv_rs_bd_to_out = ["${aci_rest.rest_l3_ext_out.id}"] # Relation to l3extOut class. Cardinality - N_TO_M 
}
