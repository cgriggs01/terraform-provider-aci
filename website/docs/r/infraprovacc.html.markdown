---
layout: "aci"
page_title: "ACI: aci_vlan_encapsulationfor_vxlan_traffic"
sidebar_current: "docs-aci-resource-vlan_encapsulationfor_vxlan_traffic"
description: |-
  Manages ACI Vlan Encapsulation for Vxlan Traffic
---

# aci_vlan_encapsulationfor_vxlan_traffic #
Manages ACI Vlan Encapsulation for Vxlan Traffic

## Example Usage ##

```hcl
resource "aci_vlan_encapsulationfor_vxlan_traffic" "example" {

  attachable_access_entity_profile_dn  = "${aci_attachable_access_entity_profile.example.id}"
  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `attachable_access_entity_profile_dn` - (Required) Distinguished name of parent AttachableAccessEntityProfile object.
* `annotation` - (Optional) annotation for object vlan_encapsulationfor_vxlan_traffic.
* `name_alias` - (Optional) name_alias for object vlan_encapsulationfor_vxlan_traffic.

* `relation_infra_rs_func_to_epg` - (Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.
                


