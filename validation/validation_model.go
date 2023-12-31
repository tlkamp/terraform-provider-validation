package validation

import (
	"errors"

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

// FromMap updates the validationDocument fields from the corresponding fields in the map.
// TODO: use mitchellh/mapstructure instead of doing this manually -- requires a refactor
func (v *validationDocument) FromMap(m map[string]any) error {
	if _, ok := m[conditionKey]; !ok {
		return errors.New("condition is required")
	}

	if _, ok := m[conditionKey].(bool); !ok {
		return errors.New("condition must be type bool")
	}

	if _, ok := m[summaryKey]; !ok {
		return errors.New("summary is required")
	}

	if _, ok := m[summaryKey].(string); !ok {
		return errors.New("summary must be type string")
	}

	v.condition = m[conditionKey].(bool)
	v.summary = m[summaryKey].(string)

	if d, ok := m[detailsKey]; ok {
		v.details = d.(string)
	}

	return nil
}
