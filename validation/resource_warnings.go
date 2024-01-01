package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWarnings() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes one or more warnings to be shown during execution if the condition is true.",
		CreateContext: resourceWarningsCreate,
		ReadContext:   basicCRUDFunc,
		UpdateContext: basicCRUDFunc,
		DeleteContext: basicCRUDFunc,
		Schema: map[string]*schema.Schema{
			"warning": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: warnResourceSchema,
				},
				Required: true,
			},
		},
	}
}

func resourceWarningsCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	warnings := data.Get("warning")

	warnList := warnings.([]interface{})

	data.SetId(uuid.NewString())

	return parseValidations(warnList, true)
}
