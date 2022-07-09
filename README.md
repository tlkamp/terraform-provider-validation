# Validation Provider
Perform extended validation on Terraform configurations, at either `plan` or `apply` time.

This aims to address a core limitation of Terraform: validating multiple variables in the same context.

## Usage
This provider requires no configuration.

```hcl
provider "validation" {}
```
