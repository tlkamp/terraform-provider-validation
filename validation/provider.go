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
	errorSchema = map[string]*schema.Schema{
		conditionKey: &schema.Schema{
			Type:        schema.TypeBool,
			Required:    true,
			Description: errorConditionDescription,
			ForceNew:    true,
		},
		summaryKey: &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: messageDescription,
		},
		detailsKey: &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: detailsDescription,
		},
	}
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
