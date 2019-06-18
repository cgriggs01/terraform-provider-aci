---
layout: "aci"
page_title: "ACI: aci_access_port_block"
sidebar_current: "docs-aci-resource-access_port_block"
description: |-
  Manages ACI Access Port Block
---

# aci_access_port_block #
Manages ACI Access Port Block

## Example Usage ##

```hcl
resource "aci_access_port_block" "example" {

  access_port_selector_dn  = "${aci_access_port_selector.example.id}"

    name  = "example"

  annotation  = "example"
  from_card  = "example"
  from_port  = "example"
  name_alias  = "example"
  to_card  = "example"
  to_port  = "example"
}
```
## Argument Reference ##
* `access_port_selector_dn` - (Required) Distinguished name of parent AccessPortSelector object.
* `name` - (Required) name of Object access_port_block.
* `annotation` - (Optional) annotation for object access_port_block.
* `from_card` - (Optional) from module
* `from_port` - (Optional) port block from port
* `name_alias` - (Optional) name_alias for object access_port_block.
* `to_card` - (Optional) port block to module
* `to_port` - (Optional) port block to port

* `relation_infra_rs_acc_bndl_subgrp` - (Optional) Relation to class infraAccBndlSubgrp. Cardinality - N_TO_ONE. Type - String.
                


