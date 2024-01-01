variable "one" {}
variable "two" {}

resource "validation_warnings" "warns" {

  warning {
    condition = var.one == var.two
    summary   = "var.one and var.two are equal. This will cause an error in future versions"
    details   = <<EOF
In a future release of this code, var.one and var.two may no longer be equal. 
Please consider modifying the values to avoid execution failures.
var.one: ${var.one}
var.two: ${var.two}
EOF
  }

  warning {
    condition = var.one % 2 == 0
    summary   = "var.one should not be even"
  }

  warning {
    condition = var.two % 2 != 0
    summary   = "var.two should not be odd"
  }
}
