---
layout: "aci"
page_title: "ACI: aci_access_port_selector"
sidebar_current: "docs-aci-resource-access_port_selector"
description: |-
  Manages ACI Access Port Selector
---

# aci_access_port_selector #
Manages ACI Access Port Selector

## Example Usage ##

```hcl
resource "aci_access_port_selector" "example" {

  leaf_interface_profile_dn  = "${aci_leaf_interface_profile.example.id}"

    name  = "example"


    access_port_selector_type  = "example"

  annotation  = "example"
  name_alias  = "example"
  access_port_selector_type  = "example"
}
```
## Argument Reference ##
* `leaf_interface_profile_dn` - (Required) Distinguished name of parent LeafInterfaceProfile object.
* `name` - (Required) name of Object access_port_selector.
* `access_port_selector_type` - (Required) access_port_selector_type of Object access_port_selector.
* `annotation` - (Optional) annotation for object access_port_selector.
* `name_alias` - (Optional) name_alias for object access_port_selector.
* `access_port_selector_type` - (Optional) host port selector type

* `relation_infra_rs_acc_base_grp` - (Optional) Relation to class infraAccBaseGrp. Cardinality - N_TO_ONE. Type - String.
                


