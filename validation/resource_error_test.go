package validation

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const errResource = "validation_error.err"

func TestAccValidationErrorBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(errResource),
		Steps: []resource.TestStep{
			{
				Config: testAccValidationErrorConfig(false, "summary message", "more details"),
				Check:  resource.ComposeTestCheckFunc(testAccCheckResourceExists(errResource)),
			},
			{
				Config:  testAccValidationErrorConfig(false, "summary message", "more details"),
				Destroy: true,
			},
		},
	})
}

func TestAccValidationError_Error(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(errResource),
		Steps: []resource.TestStep{
			{
				Config:      testAccValidationErrorConfig(true, "error happened", "details are here"),
				ExpectError: regexp.MustCompile("error happened"),
			},
		},
	})
}

func TestAccValidationErrorUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		//CheckDestroy: testAccValidationErrorDestroyed("validation_error.err"),
		Steps: []resource.TestStep{
			{ // Create resource
				Config: testAccValidationErrorConfig(false, "summary", "details"),
				Check:  testAccCheckResourceExists(errResource),
			},
			{ // Summary Change
				Config: testAccValidationErrorConfig(false, "new summary", "details"),
				Check:  testAccCheckValidationErrFields(errResource, "new summary", "details"),
			},
			{ // Details change
				Config: testAccValidationErrorConfig(false, "new summary", "new details"),
				Check:  testAccCheckValidationErrFields(errResource, "new summary", "new details"),
			},
			{ // Condition changes to true
				Config:      testAccValidationErrorConfig(true, "new summary", "new details"),
				ExpectError: regexp.MustCompile("new summary"),
				Check:       testAccCheckResourceNotExist(errResource),
			},
		},
	})
}

func testAccCheckValidationErrFields(a, sm, d string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[a]
		if !ok {
			return fmt.Errorf("resource not found: %s", a)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID set")
		}
		summary := rs.Primary.Attributes["summary"]
		details, ok := rs.Primary.Attributes["details"]

		if summary != sm {
			return fmt.Errorf("summary attribute does not match expected: %s, %s", summary, sm)
		}

		if details != d {
			return fmt.Errorf("details attribute does not match expected: %s, %s", details, d)
		}
		return nil
	}
}

func testAccValidationErrorConfig(cond bool, summary, details string) string {
	return fmt.Sprintf(`
	resource "validation_error" "err" {
        condition = %t
        summary = "%s"
        details = "%s"
    }`, cond, summary, details)
}
