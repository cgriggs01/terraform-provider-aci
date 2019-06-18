---
layout: "aci"
page_title: "ACI: aci_end_point_retention_policy"
sidebar_current: "docs-aci-resource-end_point_retention_policy"
description: |-
  Manages ACI End Point Retention Policy
---

# aci_end_point_retention_policy #
Manages ACI End Point Retention Policy

## Example Usage ##

```hcl
resource "aci_end_point_retention_policy" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  bounce_age_intvl  = "example"
  bounce_trig  = "example"
  hold_intvl  = "example"
  local_ep_age_intvl  = "example"
  move_freq  = "example"
  name_alias  = "example"
  remote_ep_age_intvl  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object end_point_retention_policy.
* `annotation` - (Optional) annotation for object end_point_retention_policy.
* `bounce_age_intvl` - (Optional) aging interval for endpoint migration
* `bounce_trig` - (Optional) bounce trigger
* `hold_intvl` - (Optional) hold interval
* `local_ep_age_intvl` - (Optional) local endpoint aging interval
* `move_freq` - (Optional) move frequency
* `name_alias` - (Optional) name_alias for object end_point_retention_policy.
* `remote_ep_age_intvl` - (Optional) remote endpoint aging interval



