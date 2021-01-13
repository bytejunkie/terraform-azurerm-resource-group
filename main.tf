terraform {
  backend "s3" {}
  required_version = ">= 0.12"
}

resource "azurerm_resource_group" "my_resource_group" {
  name     = join(var.name_separator, var.name_strings)
  location = var.location

  tags = merge({
    Module-Name = "Azure Storage Account"
    Author      = "bytejunkie - matt@bytejunkie.dev"
    },
    var.tags
  )
}
