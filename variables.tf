variable "name_strings" {
  description = "This should be a list of strings which in conjunction with the seperator make up the resource group name"
  default     = null

}
variable "name_separator" {
  description = "Used with name_strings to make up the resource group name"
  default     = null

}
variable "location" {
  description = "the region to deploy the resource to."
  default     = null

}
variable "tags" {
  default = null

}

