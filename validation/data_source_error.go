package validation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var errorDataSchema = map[string]*schema.Schema{
	conditionKey: {
		Description: errorConditionDescription,
		Type:        schema.TypeBool,
		Required:    true,
	},
	summaryKey: summarySchema,
	detailsKey: detailsSchema,
}

func dataSourceError() *schema.Resource {
	return &schema.Resource{
		Description: "Causes an error to be thrown if condition is true.",
		ReadContext: dataSourceErrorRead,
		Schema:      errorDataSchema,
	}
}

func dataSourceErrorRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	vd := &validationDocument{}

	vd.condition = data.Get(conditionKey).(bool)
	vd.summary = data.Get(summaryKey).(string)

	if ptr, ok := data.GetOk(detailsKey); ok {
		vd.details = ptr.(string)
	}

	if vd.Validate() {
		diags = append(diags, vd.Diag())
	}

	data.SetId("none")
	return diags
}
