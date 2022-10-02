package validation

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
)

func TestValidationDocument(t *testing.T) {
	vd := &validationDocument{}
	assert.Empty(t, vd)
}

func TestValidationDocument_Validate(t *testing.T) {
	vdTrue := &validationDocument{condition: true}
	vdFalse := &validationDocument{condition: false}

	assert.True(t, vdTrue.Validate())
	assert.False(t, vdFalse.Validate())
}

func TestValidationDocument_Diag(t *testing.T) {
	vd := &validationDocument{condition: true, summary: "Summary", details: "Deets"}

	d := vd.Diag()
	assert.Equal(t, "Summary", d.Summary)
	assert.Equal(t, "Deets", d.Detail)
	assert.Equal(t, diag.Error, d.Severity)

	vd.severity = diag.Warning
	d = vd.Diag()

	assert.Equal(t, diag.Warning, d.Severity)
}
