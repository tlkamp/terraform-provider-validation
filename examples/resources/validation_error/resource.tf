variable "one" {}
variable "two" {}

resource "validation_error" "error" {
  condition = var.one == var.two
  summary   = "var.one and var.two must never be equal"
  details   = <<EOF
When var.one and var.two are equal, bad things can happen.
Please use differing values for these inputs.
var.one: ${var.one}
var.two: ${var.two}
EOF
}
