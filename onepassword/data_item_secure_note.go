package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemSecureNote() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemSecureNoteRead,
		Schema: resourceItemSecureNote().Schema,
	}
}
