package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceError() *schema.Resource {
	return &schema.Resource{
		Description: "Causes an error to be thrown if condition is true.",
		ReadContext: dataSourceErrorRead,
		Schema: map[string]*schema.Schema{
			conditionKey: {
				Description: errorConditionDescription,
				Type:        schema.TypeBool,
				Required:    true,
			},
			summaryKey: summarySchema,
			detailsKey: detailsSchema,
		},
	}
}

func dataSourceErrorRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var details string
	cond := data.Get(conditionKey).(bool)
	msg := data.Get(summaryKey).(string)

	if ptr, ok := data.GetOk(detailsKey); ok {
		details = ptr.(string)
	}

	if cond {
		diags = append(diags, buildDiag(diag.Error, msg, details))
		return diags
	}

	data.SetId(uuid.NewString())
	return diags
}
