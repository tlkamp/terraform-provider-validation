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

func resourceError() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes an error to be thrown during execution if the condition is true.",
		CreateContext: resourceErrorCreate,
		ReadContext:   resourceErrorRead,
		UpdateContext: resourceErrorUpdate,
		DeleteContext: resourceErrorDelete,

		Schema: map[string]*schema.Schema{
			conditionKey: &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: errorConditionDescription,
				ForceNew:    true,
			},
			summaryKey: summarySchema,
			detailsKey: detailsSchema,
		},
	}
}

func resourceErrorCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	cond := data.Get(conditionKey).(bool)
	summary := data.Get(summaryKey).(string)
	var details string
	if ptr, ok := data.GetOk(detailsKey); ok {
		details = ptr.(string)
	}

	if cond {
		return []diag.Diagnostic{buildDiag(diag.Error, summary, details)}
	}

	id := uuid.New()
	data.SetId(id.String())

	return resourceErrorRead(ctx, data, i)
}

// not a real resource, so nothing to read.
func resourceErrorRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

// return nil because a change to condition causes force-recreate, otherwise update the changed fields and that's it.
func resourceErrorUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

// not a real resource, so delete it from state.
func resourceErrorDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}
