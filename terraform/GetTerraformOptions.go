package terraform

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetTerraformOptions(terraformDirectoryPath string, variables map[string]interface{}) *terraform.Options {

	terraformOptions := &terraform.Options{

		TerraformDir: terraformDirectoryPath,
		Vars:         variables,
		NoColor:      true,
	}

	return terraformOptions
}
