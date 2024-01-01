package validation

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceErrors() *schema.Resource {
	return &schema.Resource{
		Description:   "Causes one or more errors to be thrown during execution if the condition is true",
		CreateContext: resourceErrorsCreate,
		ReadContext:   basicCRUDFunc,
		UpdateContext: basicCRUDFunc,
		DeleteContext: basicCRUDFunc,
		Schema: map[string]*schema.Schema{
			"error": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: errorResourceSchema,
				},
				Required: true,
			},
		},
	}
}

func resourceErrorsCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	errs := data.Get("error")
	errsList := errs.([]interface{})

	for _, e := range errsList {
		v := &validationDocument{}
		eMap := e.(map[string]interface{})

		if err := v.FromMap(eMap); err != nil {
			diags = append(diags, diag.Diagnostic{
				Summary: "Error converting input to validationDocument",
				Detail:  err.Error(),
			})
		}

		if v.Validate() {
			diags = append(diags, v.Diag())
			return diags
		}
	}

	data.SetId(uuid.NewString())

	return nil
}
