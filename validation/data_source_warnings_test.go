package validation

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSrcWarns = "data.validation_warnings.warns"

func TestDataSourceWarningsBasic(t *testing.T) {
	vs := []testValidation{
		{Condition: true, Summary: "summary"},
		{Condition: true, Summary: "Summary2", Details: "with details"},
		{},
	}

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataValidationWarningsConfig(t, vs),
				Check:  resource.ComposeTestCheckFunc(testAccCheckResourceExists(dataSrcWarns)),
			},
		},
	})
}

func testAccDataValidationWarningsConfig(t *testing.T, v []testValidation) string {
	t.Helper()

	rsc := `data "validation_warnings" "warns" {
	{{- range . }}
	warning {
		condition = {{ .Condition }}
		summary   = "{{ .Summary }}"
		{{ if .Details }}details   = "{{ .Details }}"{{ end }}
	}
	{{ end }}
}
`

	tmpl := template.New("test")
	out := &bytes.Buffer{}
	parsed, _ := tmpl.Parse(rsc)

	if err := parsed.Execute(out, v); err != nil {
		t.Error(err)
	}

	return out.String()
}
