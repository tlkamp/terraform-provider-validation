package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	errorConditionDescription = "The condition which, if true, causes an error to be thrown."
	messageDescription        = "The message displayed to the user if the condition is true."
	detailsDescription        = "More details about the message being displayed to the user, if any."
)

var errorResourceSchema = map[string]*schema.Schema{
	conditionKey: {
		Type:        schema.TypeBool,
		Required:    true,
		Description: errorConditionDescription,
		ForceNew:    true,
	},
	summaryKey: summarySchema,
	detailsKey: detailsSchema,
}

func resourceError() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes an error to be thrown during execution if the condition is true.",
		CreateContext: resourceErrorCreate,
		ReadContext:   basicCRUDFunc,
		UpdateContext: basicCRUDFunc,
		DeleteContext: basicCRUDFunc,
		Schema:        errorResourceSchema,
	}
}

func resourceErrorCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	vd := &validationDocument{}

	vd.condition = data.Get(conditionKey).(bool)
	vd.summary = data.Get(summaryKey).(string)

	if ptr, ok := data.GetOk(detailsKey); ok {
		vd.details = ptr.(string)
	}

	if vd.Validate() {
		diags = append(diags, vd.Diag())
		return diags
	}

	data.SetId(uuid.NewString())

	return nil
}
