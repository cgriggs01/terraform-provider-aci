---
layout: "aci"
page_title: "ACI: aci_filter_entry"
sidebar_current: "docs-aci-resource-filter_entry"
description: |-
  Manages ACI Filter Entry
---

# aci_filter_entry #
Manages ACI Filter Entry

## Example Usage ##

```hcl
resource "aci_filter_entry" "example" {

  filter_dn  = "${aci_filter.example.id}"

    name  = "example"

  annotation  = "example"
  apply_to_frag  = "example"
  arp_opc  = "example"
  d_from_port  = "example"
  d_to_port  = "example"
  ether_t  = "example"
  icmpv4_t  = "example"
  icmpv6_t  = "example"
  match_dscp  = "example"
  name_alias  = "example"
  prot  = "example"
  s_from_port  = "example"
  s_to_port  = "example"
  stateful  = "example"
  tcp_rules  = "example"
}
```
## Argument Reference ##
* `filter_dn` - (Required) Distinguished name of parent Filter object.
* `name` - (Required) name of Object filter_entry.
* `annotation` - (Optional) annotation for object filter_entry.
* `apply_to_frag` - (Optional) fragment
* `arp_opc` - (Optional) open peripheral codes
* `d_from_port` - (Optional) end of the destination port range
* `d_to_port` - (Optional) start of the destination port range
* `ether_t` - (Optional) ethertype
* `icmpv4_t` - (Optional) 
* `icmpv6_t` - (Optional) 
* `match_dscp` - (Optional) match_dscp for object filter_entry.
* `name_alias` - (Optional) name_alias for object filter_entry.
* `prot` - (Optional) level 3 ip protocol
* `s_from_port` - (Optional) start of the source port range
* `s_to_port` - (Optional) end of the source port range
* `stateful` - (Optional) stateful entry
* `tcp_rules` - (Optional) tcp flags



