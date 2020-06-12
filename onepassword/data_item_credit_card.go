package onepassword

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceItemCreditCard() *schema.Resource {
	return &schema.Resource{
		Read:   resourceItemCreditCardRead,
		Schema: resourceItemCreditCard().Schema,
	}
}
