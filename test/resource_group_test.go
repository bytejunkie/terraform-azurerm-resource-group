package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformResourceGroup(t *testing.T) {
	t.Parallel()

	resourceGroupName := "test01-mns-rg-01"
	resourceGroupNameStrings := []string{"test01", "mns", "rg", "01"}
	resourceGroupLocation := "westeurope"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/tf-azure-resource-group",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name_strings":   resourceGroupNameStrings,
			"name_separator": "-",
		},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualResourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	actualResourceGroupLocation := terraform.Output(t, terraformOptions, "resource_group_location")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, resourceGroupName, actualResourceGroupName)
	assert.Equal(t, resourceGroupLocation, actualResourceGroupLocation)
}
func TestTerraformResourceGroupWithTags(t *testing.T) {
	t.Parallel()

	resourceGroupName := "test02-mns-rg-01"
	resourceGroupNameStrings := []string{"test02", "mns", "rg", "01"}
	resourceGroupLocation := "westeurope"

	environment := "production"
	resourceGroupTags := map[string]string{"environment": environment, "location": "westeurope"}

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/tf-azure-resource-group",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name_strings":   resourceGroupNameStrings,
			"name_separator": "-",
			"tags":           resourceGroupTags,
		},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualResourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	actualResourceGroupLocation := terraform.Output(t, terraformOptions, "resource_group_location")
	actualResourceGroupTags := terraform.OutputMap(t, terraformOptions, "resource_group_tags")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, resourceGroupName, actualResourceGroupName)
	assert.Equal(t, resourceGroupLocation, actualResourceGroupLocation)
	assert.Equal(t, resourceGroupTags, actualResourceGroupTags)
}
