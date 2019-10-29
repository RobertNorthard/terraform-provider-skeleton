package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataCoinbaseAddress() *schema.Resource {
        return &schema.Resource{
                Read:   resourceCoinbaseSourceAddressRead,
                Schema: map[string]*schema.Schema{
			"api_endpoint": &schema.Schema{
				Type:     schema.TypeString,
                                Computed: true,
			},
                },
        }
}

func resourceCoinbaseSourceAddressRead(d *schema.ResourceData, m interface{}) error {
        endpoint := "https://api.coinbase.com/v2" 
        d.SetId(endpoint)
        d.Set("api_endpoint", endpoint)
        return nil
}
