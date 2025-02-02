package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciLeafInterfaceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciLeafInterfaceProfileCreate,
		Update: resourceAciLeafInterfaceProfileUpdate,
		Read:   resourceAciLeafInterfaceProfileRead,
		Delete: resourceAciLeafInterfaceProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciLeafInterfaceProfileImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
func getRemoteLeafInterfaceProfile(client *client.Client, dn string) (*models.LeafInterfaceProfile, error) {
	infraAccPortPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	infraAccPortP := models.LeafInterfaceProfileFromContainer(infraAccPortPCont)

	if infraAccPortP.DistinguishedName == "" {
		return nil, fmt.Errorf("LeafInterfaceProfile %s not found", infraAccPortP.DistinguishedName)
	}

	return infraAccPortP, nil
}

func setLeafInterfaceProfileAttributes(infraAccPortP *models.LeafInterfaceProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(infraAccPortP.DistinguishedName)
	d.Set("description", infraAccPortP.Description)
	d.Set("name", GetMOName(infraAccPortP.DistinguishedName))
	infraAccPortPMap, _ := infraAccPortP.ToMap()

	d.Set("annotation", infraAccPortPMap["annotation"])
	d.Set("name_alias", infraAccPortPMap["nameAlias"])
	return d
}

func resourceAciLeafInterfaceProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	infraAccPortP, err := getRemoteLeafInterfaceProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLeafInterfaceProfileAttributes(infraAccPortP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciLeafInterfaceProfileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	infraAccPortPAttr := models.LeafInterfaceProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraAccPortPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraAccPortPAttr.NameAlias = NameAlias.(string)
	}
	infraAccPortP := models.NewLeafInterfaceProfile(fmt.Sprintf("infra/accportprof-%s", name), "uni", desc, infraAccPortPAttr)

	err := aciClient.Save(infraAccPortP)
	if err != nil {
		return err
	}

	d.SetId(infraAccPortP.DistinguishedName)
	return resourceAciLeafInterfaceProfileRead(d, m)
}

func resourceAciLeafInterfaceProfileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	infraAccPortPAttr := models.LeafInterfaceProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraAccPortPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraAccPortPAttr.NameAlias = NameAlias.(string)
	}
	infraAccPortP := models.NewLeafInterfaceProfile(fmt.Sprintf("infra/accportprof-%s", name), "uni", desc, infraAccPortPAttr)

	infraAccPortP.Status = "modified"

	err := aciClient.Save(infraAccPortP)

	if err != nil {
		return err
	}

	d.SetId(infraAccPortP.DistinguishedName)
	return resourceAciLeafInterfaceProfileRead(d, m)

}

func resourceAciLeafInterfaceProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	infraAccPortP, err := getRemoteLeafInterfaceProfile(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLeafInterfaceProfileAttributes(infraAccPortP, d)

	return nil
}

func resourceAciLeafInterfaceProfileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "infraAccPortP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
