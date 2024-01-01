package validation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			"validation_error":    resourceError(),
			"validation_errors":   resourceErrors(),
			"validation_warning":  resourceWarning(),
			"validation_warnings": resourceWarnings(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"validation_error":    dataSourceError(),
			"validation_errors":   dataSourceErrors(),
			"validation_warning":  dataSourceWarning(),
			"validation_warnings": dataSourceWarnings(),
		},
	}
}

func basicCRUDFunc(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func parseValidations(items []interface{}, warn bool) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, item := range items {
		v := &validationDocument{}

		if warn {
			v.severity = diag.Warning
		}

		iMap := item.(map[string]interface{})

		if err := v.FromMap(iMap); err != nil {
			diags = append(diags, diag.Diagnostic{
				Summary: "Error converting input to validationDocument",
				Detail:  err.Error(),
			})
		}

		if v.Validate() {
			diags = append(diags, v.Diag())
		}
	}

	return diags
}
