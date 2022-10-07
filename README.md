# Validation Provider
Perform extended validation on Terraform configurations, at either `plan` or `apply` time.

This aims to address a core limitation of Terraform: validating multiple variables in the same context.

## Usage
This provider requires no configuration.

### Provider

```hcl
provider "validation" {}
```

### Resources

#### `validation_error`
The `validation_error` resource can be used to evaluate values only known at apply time, and stop an
in-progress Terraform execution based upon their values.

```hcl
variable "one" {}
variable "two" {}

resource "validation_error" "error" {
  condition = var.one == var.two
  summary = "var.one and var.two must never be equal"
  details = <<EOF
When var.one and var.two are equal, bad things can happen.
Please use differing values for these inputs.
var.one: ${var.one}
var.two: ${var.two}
EOF
}
```

#### `validation_warning`
The `validation_warning` resource can be used to evaluate values only known at apply time. This resource
will print a message to the user, but not cause the Terraform execution to fail.

This can be useful for deprecation notices, etc.

```hcl
variable "one" {}
variable "two" {}

resource "validation_warning" "warn" {
  condition = var.one == var.two
  summary = "var.one and var.two are equal. This will cause an error in future versions"
  details = <<EOF
In a future release of this code, var.one and var.two may no longer be equal. Please consider modifying the values to
be distinct to avoid execution failures.
var.one: ${var.one}
var.two: ${var.two}
EOF
}
```

### Data Sources
Each resource above has an accompanying data source, which can be used to show warnings or errors during
the plan phase and may not stop an in-progress Terraform execution.
