package validation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: Refactor into a function that returns the schema
var warnDataSchema = map[string]*schema.Schema{
	conditionKey: {
		Description: warnConditionDescription,
		Type:        schema.TypeBool,
		Required:    true,
	},
	summaryKey: summarySchema,
	detailsKey: detailsSchema,
}

func dataSourceWarning() *schema.Resource {
	return &schema.Resource{
		Description: "Causes a warning message to be printed if condition is true.",
		ReadContext: dataSourceWarningRead,
		Schema:      warnDataSchema,
	}
}

func dataSourceWarningRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	vd := &validationDocument{severity: diag.Warning}

	data.SetId("none")

	vd.condition = data.Get(conditionKey).(bool)
	vd.summary = data.Get(summaryKey).(string)

	if ptr, ok := data.GetOk(detailsKey); ok {
		vd.details = ptr.(string)
	}

	if vd.Validate() {
		diags = append(diags, vd.Diag())
	}

	return diags
}
