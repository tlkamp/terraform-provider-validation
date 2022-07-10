package validation

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const warnResource = "validation_warning.warn"

func TestAccValidationWarningBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(warnResource),
		Steps: []resource.TestStep{
			{
				Config: testAccValidationWarningConfig(false, "summary message", "more details"),
				Check:  resource.ComposeTestCheckFunc(testAccCheckResourceExists(warnResource)),
			},
			{
				Config:  testAccValidationWarningConfig(false, "summary message", "more details"),
				Destroy: true,
			},
		},
	})
}

func TestAccValidationWarning(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccValidationWarningConfig(true, "warning happened", "details are here"),
				Check:  testAccCheckValidationWarnFields(warnResource, "warning happened", "details are here"),
			},
		},
	})
}

func TestAccValidationWarningUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{ // Create resource
				Config: testAccValidationWarningConfig(false, "summary", "details"),
				Check:  testAccCheckResourceExists(warnResource),
			},
			{ // Summary change
				Config: testAccValidationWarningConfig(false, "new summary", "details"),
				Check:  testAccCheckValidationErrFields(warnResource, "new summary", "details"),
			},
			{ // Details change
				Config: testAccValidationWarningConfig(false, "new summary", "new details"),
				Check:  testAccCheckValidationErrFields(warnResource, "new summary", "new details"),
			},
			{ // Condition changes to true
				Config: testAccValidationWarningConfig(true, "new summary", "new details"),
				Check:  testAccCheckResourceExists(warnResource),
			},
		},
	})
}

func testAccCheckValidationWarnFields(a, sm, d string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[a]
		if !ok {
			return fmt.Errorf("resource not found: %s", a)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID set")
		}
		summary := rs.Primary.Attributes["summary"]
		details := rs.Primary.Attributes["details"]

		if summary != sm {
			return fmt.Errorf("summary attribute does not match expected: %s, %s", summary, sm)
		}

		if details != d {
			return fmt.Errorf("details attribute does not match expected: %s, %s", details, d)
		}
		return nil
	}
}

func testAccValidationWarningConfig(cond bool, summary, details string) string {
	return fmt.Sprintf(`
	resource "validation_warning" "warn" {
        condition = %t
        summary = "%s"
        details = "%s"
    }`, cond, summary, details)
}
