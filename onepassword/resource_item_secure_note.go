package onepassword

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceItemSecureNote() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemSecureNoteRead,
		Create: resourceItemSecureNoteCreate,
		Delete: resourceItemDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				err := resourceItemSecureNoteRead(d, meta)
				return []*schema.ResourceData{d}, err
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vault": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"section": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     sectionSchema(),
			},
		},
	}
}

func resourceItemSecureNoteRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	vaultID := d.Get("vault").(string)
	v, err := m.onePassClient.ReadItem(getID(d), vaultID)
	if err != nil {
		return err
	}
	if v.Template != Category2Template(SecureNoteCategory) {
		return errors.New("item is not from " + string(SecureNoteCategory))
	}

	d.SetId(v.UUID)
	if err := d.Set("name", v.Overview.Title); err != nil {
		return err
	}
	if err := d.Set("tags", v.Overview.Tags); err != nil {
		return err
	}
	if err := d.Set("vault", v.Vault); err != nil {
		return err
	}
	if err := d.Set("notes", v.Details.Notes); err != nil {
		return err
	}
	return d.Set("section", ProcessSections(v.Details.Sections))
}

func resourceItemSecureNoteCreate(d *schema.ResourceData, meta interface{}) error {
	item := &Item{
		Vault:    d.Get("vault").(string),
		Template: Category2Template(SecureNoteCategory),
		Details: Details{
			Notes:    d.Get("notes").(string),
			Sections: ParseSections(d),
		},
		Overview: Overview{
			Title: d.Get("name").(string),
			Tags:  ParseTags(d),
		},
	}
	m := meta.(*Meta)
	err := m.onePassClient.CreateItem(item)
	if err != nil {
		return err
	}
	d.SetId(item.UUID)
	return nil
}
