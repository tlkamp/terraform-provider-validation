package validation

import (
	"bytes"
	"regexp"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSrcErrs = "data.validation_errors.errs"

func TestAccDataSourceErrorsBasic(t *testing.T) {
	vs := []testValidation{
		{Condition: false, Summary: "test data source"},
		{Condition: false, Summary: "test data source2", Details: "with details"},
		{},
	}

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccValidationDataErrorsConfig(t, vs),
				Check:  testAccCheckResourceExists(dataSrcErrs),
			},
		},
	})
}

func TestAccDataSourceErrors(t *testing.T) {
	vs := []testValidation{
		{Condition: false, Summary: "test data source"},
		{Condition: true, Summary: "test data source2", Details: "with details"},
		{},
	}
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccValidationDataErrorsConfig(t, vs),
				Check:       testAccCheckResourceNotExist(dataSrcErrs),
				ExpectError: regexp.MustCompile("test data source2"),
			},
		},
	})
}

func testAccValidationDataErrorsConfig(t *testing.T, v []testValidation) string {
	t.Helper()

	rsc := `data "validation_errors" "errs" {
	{{- range . }}
	error {
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
