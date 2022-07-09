package validation

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
)

func TestBuildDiag(t *testing.T) {
	expectedSummary := "summary"
	expectedDetails := "details"

	diagErr := buildDiag(diag.Error, expectedSummary, expectedDetails)
	assert.Equal(t, expectedSummary, diagErr.Summary)
	assert.Equal(t, expectedDetails, diagErr.Detail)
	assert.Equal(t, diag.Error, diagErr.Severity)

	diagWarn := buildDiag(diag.Warning, "diff", "diff2")
	assert.Equal(t, "diff", diagWarn.Summary)
	assert.Equal(t, "diff2", diagWarn.Detail)
	assert.Equal(t, diag.Warning, diagWarn.Severity)
}
