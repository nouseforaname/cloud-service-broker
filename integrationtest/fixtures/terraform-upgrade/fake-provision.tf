terraform {
  required_providers {
    random = {
      source  = "registry.terraform.io/hashicorp/random"
    }
  }
}

resource "random_integer" "priority" {
  min = 1
  max = 2
}

output "provision_output" {
  value = random_integer.priority.result
}