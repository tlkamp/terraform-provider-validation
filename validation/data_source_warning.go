package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWarning() *schema.Resource {
	return &schema.Resource{
		Description: "Causes a warning message to be printed if condition is true.",
		ReadContext: dataSourceWarningRead,
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

func dataSourceWarningRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var details string

	// We always create this resource
	data.SetId(uuid.NewString())

	cond := data.Get(conditionKey).(bool)
	msg := data.Get(summaryKey).(string)
	if ptr, ok := data.GetOk(detailsKey); ok {
		details = ptr.(string)
	}

	if cond {
		// But we only show a warning if the condition is true
		diags = append(diags, buildDiag(diag.Warning, msg, details))
	}

	return diags
}
