package validation

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// These are used in other tests
var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"validation": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	assert.NoError(t, Provider().InternalValidate())
}

func TestProviderImpl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func TestBasicCRUDFunc(t *testing.T) {
	assert.Nil(t, basicCRUDFunc(context.TODO(), nil, nil))
}

func TestParseValidations(t *testing.T) {
	t.Run("Warnings", func(t *testing.T) {
		var warnings []interface{}

		d := map[string]interface{}{
			"condition": true,
			"summary":   "summary",
			"details":   "deeeeets",
		}

		d2 := map[string]interface{}{
			"condition": false,
			"summary":   "s",
		}

		warnings = append(warnings, d, d2)

		diags := parseValidations(warnings, true)

		assert.NotNil(t, diags)
		assert.NotEmpty(t, diags)
		assert.Len(t, diags, 1)

		// Everything returned should be warning level
		for _, w := range diags {
			assert.Equal(t, w.Severity, diag.Warning)
		}
	})

	t.Run("Errors", func(t *testing.T) {
		var errs []interface{}

		d := map[string]interface{}{
			"condition": true,
			"summary":   "summary",
			"details":   "deeeeets",
		}

		d1 := map[string]interface{}{
			"condition": true,
			"summary":   "summary2",
			"details":   "deets",
		}

		d2 := map[string]interface{}{
			"condition": false,
			"summary":   "s",
		}

		errs = append(errs, d, d1, d2)

		diags := parseValidations(errs, false)

		assert.NotNil(t, diags)
		assert.NotEmpty(t, diags)
		assert.Len(t, diags, 2)

		// Everything returned should be error level
		for _, w := range diags {
			assert.Equal(t, w.Severity, diag.Error)
		}
	})

}

// This function is used between Resource tests
func testAccCheckResourceExists(addr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[addr]
		if !ok {
			return fmt.Errorf("resource not found: %s", addr)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource id set")
		}
		return nil
	}
}

// This function is used between Resource tests
func testAccCheckResourceNotExist(addr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[addr]
		if ok {
			return fmt.Errorf("resource still exists: %s", addr)
		}

		return nil
	}
}
