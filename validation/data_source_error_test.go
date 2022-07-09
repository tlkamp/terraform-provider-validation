package validation

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSrcErr = "data.validation_error.err"

func TestAccDataSourceErrorBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataValidationErrConfig(false, "A summary", "details"),
				Check:  testAccCheckResourceExists(dataSrcErr),
			},
		},
	})
}

func TestAccDataSourceError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataValidationErrConfig(true, "A summary", "details"),
				Check:       testAccCheckResourceNotExist(dataSrcErr),
				ExpectError: regexp.MustCompile("A summary"),
			},
		},
	})
}

func testAccDataValidationErrConfig(cond bool, summary, details string) string {
	return fmt.Sprintf(`
	data "validation_error" "err" {
        condition = %t
        summary = "%s"
		details = "%s"
    }`, cond, summary, details)
}
