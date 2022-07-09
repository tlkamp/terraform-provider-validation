package validation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func buildDiag(sev diag.Severity, summary, details string) diag.Diagnostic {
	return diag.Diagnostic{
		Severity: sev,
		Summary:  summary,
		Detail:   details,
	}
}
