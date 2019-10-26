package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSkeleton() *schema.Resource {
        return &schema.Resource{
                Create: resourceSkeletonCreate,
                Read:   resourceSkeletonRead,
                Update: resourceSkeletonUpdate,
                Delete: resourceSkeletonDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceSkeletonCreate(d *schema.ResourceData, m interface{}) error {
	return resourceSkeletonRead(d, m)
}

func resourceSkeletonRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSkeletonUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSkeletonRead(d, m)
}

func resourceSkeletonDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}