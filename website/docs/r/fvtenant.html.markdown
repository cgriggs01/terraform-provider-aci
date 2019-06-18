---
layout: "aci"
page_title: "ACI: aci_tenant"
sidebar_current: "docs-aci-resource-tenant"
description: |-
  Manages ACI Tenant
---

# aci_tenant #
Manages ACI Tenant

## Example Usage ##

```hcl
resource "aci_tenant" "example" {


    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object tenant.
* `annotation` - (Optional) annotation for object tenant.
* `name_alias` - (Optional) name_alias for object tenant.

* `relation_fv_rs_tn_deny_rule` - (Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_tenant_mon_pol` - (Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.
                


