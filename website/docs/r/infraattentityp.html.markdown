---
layout: "aci"
page_title: "ACI: aci_attachable_access_entity_profile"
sidebar_current: "docs-aci-resource-attachable_access_entity_profile"
description: |-
  Manages ACI Attachable Access Entity Profile
---

# aci_attachable_access_entity_profile #
Manages ACI Attachable Access Entity Profile

## Example Usage ##

```hcl
resource "aci_attachable_access_entity_profile" "example" {


    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object attachable_access_entity_profile.
* `annotation` - (Optional) annotation for object attachable_access_entity_profile.
* `name_alias` - (Optional) name_alias for object attachable_access_entity_profile.

* `relation_infra_rs_dom_p` - (Optional) Relation to class infraADomP. Cardinality - N_TO_M. Type - Set of String.
                


