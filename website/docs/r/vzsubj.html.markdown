---
layout: "aci"
page_title: "ACI: aci_contract_subject"
sidebar_current: "docs-aci-resource-contract_subject"
description: |-
  Manages ACI Contract Subject
---

# aci_contract_subject #
Manages ACI Contract Subject

## Example Usage ##

```hcl
resource "aci_contract_subject" "example" {

  contract_dn  = "${aci_contract.example.id}"

    name  = "example"

  annotation  = "example"
  cons_match_t  = "example"
  name_alias  = "example"
  prio  = "example"
  prov_match_t  = "example"
  rev_flt_ports  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `contract_dn` - (Required) Distinguished name of parent Contract object.
* `name` - (Required) name of Object contract_subject.
* `annotation` - (Optional) annotation for object contract_subject.
* `cons_match_t` - (Optional) consumer subject match criteria
* `name_alias` - (Optional) name_alias for object contract_subject.
* `prio` - (Optional) priority level specifier
* `prov_match_t` - (Optional) consumer subject match criteria
* `rev_flt_ports` - (Optional) enables filter to apply on ingress and egress traffic
* `target_dscp` - (Optional) target dscp

* `relation_vz_rs_subj_graph_att` - (Optional) Relation to class vnsAbsGraph. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vz_rs_sdwan_pol` - (Optional) Relation to class extdevSDWanSlaPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vz_rs_subj_filt_att` - (Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.
                


