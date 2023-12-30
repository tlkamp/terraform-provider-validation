variable "one" {}
variable "two" {}

data "validation_warning" "warn" {
  condition = var.one == var.two
  summary   = "var.one and var.two are equal. This will cause an error in future versions"
  details   = <<EOF
In a future release of this code, var.one and var.two may no longer be equal. 
Please consider modifying the values to avoid execution failures.
var.one: ${var.one}
var.two: ${var.two}
EOF
}
