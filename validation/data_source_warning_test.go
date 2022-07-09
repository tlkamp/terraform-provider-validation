package validation

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSrcWarning = "data.validation_warning.warn"

func TestAccDataSourceWarningBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataValidationWarningConfig(false, "A summary", "deets"),
				Check:  testAccCheckResourceExists(dataSrcWarning),
			},
		},
	})
}

func TestAccDataSourceWarning(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataValidationWarningConfig(true, "Other Summary", "details"),
				Check:  testAccCheckResourceExists(dataSrcWarning),
			},
		},
	})
}

func testAccDataValidationWarningConfig(cond bool, summary, details string) string {
	return fmt.Sprintf(`
	data "validation_warning" "warn" {
        condition = %t
        summary = "%s"
		details = "%s"
    }`, cond, summary, details)
}
