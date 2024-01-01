package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	warnConditionDescription = "The condition which, if true, causes a warning message to be printed."
)

var warnResourceSchema = map[string]*schema.Schema{
	conditionKey: {
		Type:        schema.TypeBool,
		Required:    true,
		Description: warnConditionDescription,
		ForceNew:    true,
	},
	summaryKey: summarySchema,
	detailsKey: detailsSchema,
}

func resourceWarning() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes a warning to be printed if the condition is true.",
		CreateContext: resourceWarningCreate,
		ReadContext:   basicCRUDFunc,
		UpdateContext: basicCRUDFunc,
		DeleteContext: basicCRUDFunc,
		Schema:        warnResourceSchema,
	}
}

func resourceWarningCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	vd := &validationDocument{severity: diag.Warning}

	vd.condition = data.Get(conditionKey).(bool)
	vd.summary = data.Get(summaryKey).(string)

	if ptr, ok := data.GetOk(detailsKey); ok {
		vd.details = ptr.(string)
	}

	if vd.Validate() {
		diags = append(diags, vd.Diag())
	}

	data.SetId(uuid.NewString())

	return diags
}
