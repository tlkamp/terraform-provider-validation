package validation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWarnings() *schema.Resource {
	return &schema.Resource{
		Description: "Causes one or more warnings to be shown if the condition is true.",
		ReadContext: dataSourceWarningsRead,
		Schema: map[string]*schema.Schema{
			"warning": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: warnDataSchema,
				},
				Required: true,
			},
		},
	}
}

func dataSourceWarningsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	data.SetId("none")
	warnings := data.Get("warning")
	warnList := warnings.([]interface{})
	return parseValidations(warnList, true)
}
