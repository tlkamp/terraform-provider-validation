package validation

import (
	"fmt"
	"testing"

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
