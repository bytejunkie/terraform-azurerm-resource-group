resource "random_pet" "pet" {
  keepers = {
    # Generate a new pet name each time we switch to a new AMI id
  }
  separator = ""

}

module "resourcegroup" {
  source = "../../"
  # insert the required variables here

  name_strings   = ["byt", "rsg", random_pet.pet.id]
  name_separator = "-"
  location       = "west europe"

  tags = {
    location = "west europe"
  }
}



