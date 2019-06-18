package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciContractSubject() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciContractSubjectRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"contract_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciContractSubjectRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("subj-%s", name)
	ContractDn := d.Get("contract_dn").(string)

	dn := fmt.Sprintf("%s/%s", ContractDn, rn)

	vzSubj, err := getRemoteContractSubject(aciClient, dn)

	if err != nil {
		return err
	}
	setContractSubjectAttributes(vzSubj, d)
	return nil
}
