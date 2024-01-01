variable "one" {}
variable "two" {}

resource "validation_errors" "errs" {

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
