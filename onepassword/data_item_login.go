package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemLogin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemLoginRead,
		Schema: resourceItemLogin().Schema,
	}
}
