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

func resourceWarning() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes a warning to be printed if the condition is true.",
		CreateContext: resourceWarningCreate,
		ReadContext:   resourceWarningRead,
		UpdateContext: resourceWarningUpdate,
		DeleteContext: resourceWarningDelete,

		Schema: map[string]*schema.Schema{
			conditionKey: &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: warnConditionDescription,
				ForceNew:    true,
			},
			summaryKey: summarySchema,
			detailsKey: detailsSchema,
		},
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

	id := uuid.New()
	data.SetId(id.String())

	return diags
}

func resourceWarningRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func resourceWarningUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return resourceWarningRead(ctx, data, i)
}

func resourceWarningDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}
