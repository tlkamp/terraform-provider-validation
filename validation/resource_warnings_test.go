package validation

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const warnsResource = "validation_warnings.warns"

func TestAccValidationWarningsBasic(t *testing.T) {
	vs := []testValidation{
		{Condition: true, Summary: "summary"},
		{Condition: true, Summary: "Summary2", Details: "with details"},
		{},
	}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(warnsResource),
		Steps: []resource.TestStep{
			{
				Config: testAccValidationWarningsConfig(t, vs),
				Check:  resource.ComposeTestCheckFunc(testAccCheckResourceExists(warnsResource)),
			},
			{
				Config: testAccValidationWarningsConfig(t, vs[:2]),
				Check:  testAccCheckResourceExists(warnsResource),
			},
			{
				Config:  testAccValidationWarningsConfig(t, vs[:2]),
				Destroy: true,
			},
		},
	})
}

func testAccValidationWarningsConfig(t *testing.T, v []testValidation) string {
	t.Helper()

	rsc := `resource "validation_warnings" "warns" {
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
