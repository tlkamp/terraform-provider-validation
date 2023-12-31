package validation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	conditionKey = "condition"
	summaryKey   = "summary"
	detailsKey   = "details"
)

var (
	summarySchema = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: messageDescription,
	}

	detailsSchema = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: detailsDescription,
	}
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"validation_error":   resourceError(),
			"validation_errors":  resourceErrors(),
			"validation_warning": resourceWarning(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"validation_error":   dataSourceError(),
			"validation_errors":  dataSourceErrors(),
			"validation_warning": dataSourceWarning(),
		},
	}
}
