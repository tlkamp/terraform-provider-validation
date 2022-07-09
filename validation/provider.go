package validation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	conditionKey = "condition"
	summaryKey   = "summary"
	detailsKey   = "details"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"validation_error":   resourceError(),
			"validation_warning": resourceWarning(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"validation_error":   dataSourceError(),
			"validation_warning": dataSourceWarning(),
		},
	}
}
