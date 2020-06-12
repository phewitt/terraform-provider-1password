package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceVault() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVaultRead,
		Schema: resourceVault().Schema,
	}
}
