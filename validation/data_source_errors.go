package validation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceErrors() *schema.Resource {
	return &schema.Resource{
		Description: "Causes an error to be thrown if the condition is true.",
		ReadContext: dataSourceErrorsRead,
		Schema: map[string]*schema.Schema{
			"error": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: errorResourceSchema,
				},
				Required: true,
			},
		},
	}
}

func dataSourceErrorsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// TODO: encapsulate this into a function that returns diags.
	errs := data.Get("error")
	errsList := errs.([]interface{})

	for _, e := range errsList {
		v := &validationDocument{}
		eMap := e.(map[string]interface{})

		if err := v.FromMap(eMap); err != nil {
			diags = append(diags, diag.Diagnostic{
				Summary: "Error converting input to validationDocument",
				Detail:  err.Error(),
			})
		}

		if v.Validate() {
			diags = append(diags, v.Diag())
		}
	}
	// TODO: end todo

	data.SetId("none")

	return diags
}
