---
layout: "aci"
page_title: "ACI: aci_lacp_policy"
sidebar_current: "docs-aci-resource-lacp_policy"
description: |-
  Manages ACI LACP Policy
---

# aci_lacp_policy #
Manages ACI LACP Policy

## Example Usage ##

```hcl
resource "aci_lacp_policy" "example" {


    name  = "example"

  annotation  = "example"
  ctrl  = "example"
  max_links  = "example"
  min_links  = "example"
  mode  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object lacp_policy.
* `annotation` - (Optional) annotation for object lacp_policy.
* `ctrl` - (Optional) LAG control properties
* `max_links` - (Optional) maximum number of links
* `min_links` - (Optional) minimum number of links in port channel
* `mode` - (Optional) policy mode
* `name_alias` - (Optional) name_alias for object lacp_policy.



