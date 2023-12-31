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

func TestFromMap(t *testing.T) {

	t.Run("Basic", func(t *testing.T) {
		m := make(map[string]interface{})

		m[conditionKey] = true
		m[summaryKey] = "test summary"
		m[detailsKey] = "some details for you here"

		v := &validationDocument{}

		assert.NoError(t, v.FromMap(m))
		assert.NotEmpty(t, v.details)
		assert.Empty(t, v.severity)
	})

	t.Run("NoCondition", func(t *testing.T) {
		m := make(map[string]interface{})

		m[summaryKey] = "test summary"

		v := &validationDocument{}

		assert.Error(t, v.FromMap(m))

		assert.Empty(t, v.condition)
		assert.Empty(t, v.summary)
		assert.Empty(t, v.details)
		assert.Empty(t, v.severity)
	})

	t.Run("ConditionWrongType", func(t *testing.T) {
		m := make(map[string]interface{})

		m[conditionKey] = "test summary"

		v := &validationDocument{}

		assert.Error(t, v.FromMap(m))

		assert.Empty(t, v.condition)
		assert.Empty(t, v.summary)
		assert.Empty(t, v.details)
		assert.Empty(t, v.severity)
	})

	t.Run("NoSummary", func(t *testing.T) {
		m := make(map[string]interface{})

		m[conditionKey] = true

		v := &validationDocument{}

		assert.Error(t, v.FromMap(m))
		assert.Empty(t, v.condition)
		assert.Empty(t, v.details)
		assert.Empty(t, v.severity)
	})

	t.Run("SummaryWrongType", func(t *testing.T) {
		m := make(map[string]interface{})

		m[conditionKey] = true
		m[summaryKey] = 1234

		v := &validationDocument{}

		assert.Error(t, v.FromMap(m))
		assert.Empty(t, v.condition)
		assert.Empty(t, v.details)
		assert.Empty(t, v.severity)

	})

}
