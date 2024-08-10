---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "validation_errors Data Source - terraform-provider-validation"
subcategory: ""
description: |-
  Causes an error to be thrown if the condition is true.
---

# validation_errors (Data Source)

Causes an error to be thrown if the condition is true.

## Example Usage

```terraform
variable "one" {}
variable "two" {}

data "validation_errors" "errs" {

  error {
    condition = var.one == var.two
    summary   = "var.one and var.two must never be equal"
    details   = <<EOF
When var.one and var.two are equal, bad things can happen.
Please use differing values for these inputs.
var.one: ${var.one}
var.two: ${var.two}
EOF
  }

  error {
    condition = var.one % 2 == 0
    summary   = "var.one cannot be even"
  }

  error {
    condition = var.two % 2 != 0
    summary   = "var.two cannot be odd"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `error` (Block List, Min: 1) (see [below for nested schema](#nestedblock--error))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--error"></a>
### Nested Schema for `error`

Required:

- `condition` (Boolean) The condition which, if true, causes an error to be thrown.
- `summary` (String) The message displayed to the user if the condition is true.

Optional:

- `details` (String) More details about the message being displayed to the user, if any.