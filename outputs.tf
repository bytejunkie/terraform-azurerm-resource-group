output "resource_group_name" {
  value = azurerm_resource_group.my_resource_group.name
}

output "resource_group_location" {
  value = azurerm_resource_group.my_resource_group.location
}

output "resource_group_tags" {
  value = azurerm_resource_group.my_resource_group.tags
}
