package validation

import (
	"bytes"
	"regexp"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const errsResource = "validation_errors.errs"

// Temporary structure until validationDocument is refactored
// TODO: use refactored validationDocument
type testValidation struct {
	Condition bool
	Summary   string
	Details   string
}

func TestAccValidationErrorsBasic(t *testing.T) {
	vs := []testValidation{
		{Condition: false, Summary: "summary"},
		{Condition: false, Summary: "Summary2", Details: "with details"},
		{},
	}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(errsResource),
		Steps: []resource.TestStep{
			{
				Config: testAccValidationErrorsConfig(t, vs),
				Check:  resource.ComposeTestCheckFunc(testAccCheckResourceExists(errsResource)),
			},
			{
				Config:  testAccValidationErrorsConfig(t, vs),
				Destroy: true,
			},
		},
	})
}

func TestAccValidationErrors_Error(t *testing.T) {
	vs := []testValidation{
		{Condition: false, Summary: "summary"},
		{Condition: true, Summary: "Summary2", Details: "with details"},
	}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceNotExist(errsResource),
		Steps: []resource.TestStep{
			{
				Config:      testAccValidationErrorsConfig(t, vs),
				ExpectError: regexp.MustCompile("Summary2"),
			},
		},
	})
}

func TestAccValidationErrorsUpdate(t *testing.T) {
	vs1 := []testValidation{
		{Condition: false, Summary: "summary"},
		{Condition: false, Summary: "Summary2", Details: "with details"},
	}

	vs2 := []testValidation{
		{Condition: false, Summary: "summary"},
		{Condition: false, Summary: "Summary different", Details: "with details"},
	}

	vs3 := []testValidation{
		{Condition: false, Summary: "summary"},
		{Condition: true, Summary: "Summary different", Details: "with details"},
	}

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{ // Create resource
				Config: testAccValidationErrorsConfig(t, vs1),
				Check:  testAccCheckResourceExists(errsResource),
			},
			{ // Update resource
				Config: testAccValidationErrorsConfig(t, vs2),
				Check:  testAccCheckResourceExists(errsResource),
			},
			{ // Condition changes to true
				Config:      testAccValidationErrorsConfig(t, vs3),
				ExpectError: regexp.MustCompile("Summary different"),
				Check:       testAccCheckResourceNotExist(errsResource),
			},
		},
	})
}

func testAccValidationErrorsConfig(t *testing.T, v []testValidation) string {
	t.Helper()

	rsc := `resource "validation_errors" "errs" {
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
