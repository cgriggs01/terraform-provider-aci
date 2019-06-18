package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciLLDPInterfacePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciLLDPInterfacePolicyCreate,
		Update: resourceAciLLDPInterfacePolicyUpdate,
		Read:   resourceAciLLDPInterfacePolicyRead,
		Delete: resourceAciLLDPInterfacePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciLLDPInterfacePolicyImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_rx_st": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"admin_tx_st": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}
func getRemoteLLDPInterfacePolicy(client *client.Client, dn string) (*models.LLDPInterfacePolicy, error) {
	lldpIfPolCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lldpIfPol := models.LLDPInterfacePolicyFromContainer(lldpIfPolCont)

	if lldpIfPol.DistinguishedName == "" {
		return nil, fmt.Errorf("LLDPInterfacePolicy %s not found", lldpIfPol.DistinguishedName)
	}

	return lldpIfPol, nil
}

func setLLDPInterfacePolicyAttributes(lldpIfPol *models.LLDPInterfacePolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lldpIfPol.DistinguishedName)
	d.Set("description", lldpIfPol.Description)
	d.Set("name", GetMOName(lldpIfPol.DistinguishedName))
	lldpIfPolMap, _ := lldpIfPol.ToMap()

	d.Set("admin_rx_st", lldpIfPolMap["adminRxSt"])
	d.Set("admin_tx_st", lldpIfPolMap["adminTxSt"])
	d.Set("annotation", lldpIfPolMap["annotation"])
	d.Set("name_alias", lldpIfPolMap["nameAlias"])
	return d
}

func resourceAciLLDPInterfacePolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	lldpIfPol, err := getRemoteLLDPInterfacePolicy(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLLDPInterfacePolicyAttributes(lldpIfPol, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciLLDPInterfacePolicyCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	lldpIfPolAttr := models.LLDPInterfacePolicyAttributes{}
	if AdminRxSt, ok := d.GetOk("admin_rx_st"); ok {
		lldpIfPolAttr.AdminRxSt = AdminRxSt.(string)
	}
	if AdminTxSt, ok := d.GetOk("admin_tx_st"); ok {
		lldpIfPolAttr.AdminTxSt = AdminTxSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		lldpIfPolAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		lldpIfPolAttr.NameAlias = NameAlias.(string)
	}
	lldpIfPol := models.NewLLDPInterfacePolicy(fmt.Sprintf("infra/lldpIfP-%s", name), "uni", desc, lldpIfPolAttr)

	err := aciClient.Save(lldpIfPol)
	if err != nil {
		return err
	}

	d.SetId(lldpIfPol.DistinguishedName)
	return resourceAciLLDPInterfacePolicyRead(d, m)
}

func resourceAciLLDPInterfacePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	lldpIfPolAttr := models.LLDPInterfacePolicyAttributes{}
	if AdminRxSt, ok := d.GetOk("admin_rx_st"); ok {
		lldpIfPolAttr.AdminRxSt = AdminRxSt.(string)
	}
	if AdminTxSt, ok := d.GetOk("admin_tx_st"); ok {
		lldpIfPolAttr.AdminTxSt = AdminTxSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		lldpIfPolAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		lldpIfPolAttr.NameAlias = NameAlias.(string)
	}
	lldpIfPol := models.NewLLDPInterfacePolicy(fmt.Sprintf("infra/lldpIfP-%s", name), "uni", desc, lldpIfPolAttr)

	lldpIfPol.Status = "modified"

	err := aciClient.Save(lldpIfPol)

	if err != nil {
		return err
	}

	d.SetId(lldpIfPol.DistinguishedName)
	return resourceAciLLDPInterfacePolicyRead(d, m)

}

func resourceAciLLDPInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	lldpIfPol, err := getRemoteLLDPInterfacePolicy(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLLDPInterfacePolicyAttributes(lldpIfPol, d)

	return nil
}

func resourceAciLLDPInterfacePolicyDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "lldpIfPol")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
