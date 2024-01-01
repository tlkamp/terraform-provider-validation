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
	errs := data.Get("error")
	errsList := errs.([]interface{})

	data.SetId("none")

	return parseValidations(errsList, false)
}
