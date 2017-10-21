package opentelekomcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// FAIL, dangling resource
func TestAccComputeV2FloatingIP_importBasic(t *testing.T) {
	resourceName := "opentelekomcloud_compute_floatingip_v2.fip_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2FloatingIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2FloatingIP_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
