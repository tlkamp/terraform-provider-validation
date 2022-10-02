package validation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// validationDocument is a logical representation of a validation resource
type validationDocument struct {
	condition bool
	summary   string
	details   string
	severity  diag.Severity // Default is Error (0)
}

// Validate returns the result of the condition
func (v *validationDocument) Validate() bool {
	return v.condition
}

// Diag returns the diag.Diagnostic representation of the validationDocument
func (v *validationDocument) Diag() diag.Diagnostic {
	return diag.Diagnostic{
		Severity: v.severity,
		Summary:  v.summary,
		Detail:   v.details,
	}
}
