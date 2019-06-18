---
layout: "aci"
page_title: "ACI: aci_port_security_policy"
sidebar_current: "docs-aci-resource-port_security_policy"
description: |-
  Manages ACI Port Security Policy
---

# aci_port_security_policy #
Manages ACI Port Security Policy

## Example Usage ##

```hcl
resource "aci_port_security_policy" "example" {


    name  = "example"

  annotation  = "example"
  maximum  = "example"
  mode  = "example"
  name_alias  = "example"
  timeout  = "example"
  violation  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object port_security_policy.
* `annotation` - (Optional) annotation for object port_security_policy.
* `maximum` - (Optional) maximum for object port_security_policy.
* `mode` - (Optional) bgp domain mode
* `name_alias` - (Optional) name_alias for object port_security_policy.
* `timeout` - (Optional) amount of time between authentication attempts
* `violation` - (Optional) violation for object port_security_policy.



