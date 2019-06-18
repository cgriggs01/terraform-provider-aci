package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciMiscablingProtocolInterfacePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciMiscablingProtocolInterfacePolicyCreate,
		Update: resourceAciMiscablingProtocolInterfacePolicyUpdate,
		Read:   resourceAciMiscablingProtocolInterfacePolicyRead,
		Delete: resourceAciMiscablingProtocolInterfacePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciMiscablingProtocolInterfacePolicyImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_st": &schema.Schema{
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
func getRemoteMiscablingProtocolInterfacePolicy(client *client.Client, dn string) (*models.MiscablingProtocolInterfacePolicy, error) {
	mcpIfPolCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	mcpIfPol := models.MiscablingProtocolInterfacePolicyFromContainer(mcpIfPolCont)

	if mcpIfPol.DistinguishedName == "" {
		return nil, fmt.Errorf("MiscablingProtocolInterfacePolicy %s not found", mcpIfPol.DistinguishedName)
	}

	return mcpIfPol, nil
}

func setMiscablingProtocolInterfacePolicyAttributes(mcpIfPol *models.MiscablingProtocolInterfacePolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(mcpIfPol.DistinguishedName)
	d.Set("description", mcpIfPol.Description)
	d.Set("name", GetMOName(mcpIfPol.DistinguishedName))
	mcpIfPolMap, _ := mcpIfPol.ToMap()

	d.Set("admin_st", mcpIfPolMap["adminSt"])
	d.Set("annotation", mcpIfPolMap["annotation"])
	d.Set("name_alias", mcpIfPolMap["nameAlias"])
	return d
}

func resourceAciMiscablingProtocolInterfacePolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	mcpIfPol, err := getRemoteMiscablingProtocolInterfacePolicy(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setMiscablingProtocolInterfacePolicyAttributes(mcpIfPol, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciMiscablingProtocolInterfacePolicyCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	mcpIfPolAttr := models.MiscablingProtocolInterfacePolicyAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		mcpIfPolAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		mcpIfPolAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		mcpIfPolAttr.NameAlias = NameAlias.(string)
	}
	mcpIfPol := models.NewMiscablingProtocolInterfacePolicy(fmt.Sprintf("infra/mcpIfP-%s", name), "uni", desc, mcpIfPolAttr)

	err := aciClient.Save(mcpIfPol)
	if err != nil {
		return err
	}

	d.SetId(mcpIfPol.DistinguishedName)
	return resourceAciMiscablingProtocolInterfacePolicyRead(d, m)
}

func resourceAciMiscablingProtocolInterfacePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	mcpIfPolAttr := models.MiscablingProtocolInterfacePolicyAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		mcpIfPolAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		mcpIfPolAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		mcpIfPolAttr.NameAlias = NameAlias.(string)
	}
	mcpIfPol := models.NewMiscablingProtocolInterfacePolicy(fmt.Sprintf("infra/mcpIfP-%s", name), "uni", desc, mcpIfPolAttr)

	mcpIfPol.Status = "modified"

	err := aciClient.Save(mcpIfPol)

	if err != nil {
		return err
	}

	d.SetId(mcpIfPol.DistinguishedName)
	return resourceAciMiscablingProtocolInterfacePolicyRead(d, m)

}

func resourceAciMiscablingProtocolInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	mcpIfPol, err := getRemoteMiscablingProtocolInterfacePolicy(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setMiscablingProtocolInterfacePolicyAttributes(mcpIfPol, d)

	return nil
}

func resourceAciMiscablingProtocolInterfacePolicyDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "mcpIfPol")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
