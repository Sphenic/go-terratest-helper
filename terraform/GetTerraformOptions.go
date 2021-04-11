package terraform

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetTerraformOptions(terraformDirectoryPath string) *terraform.Options {

	terraformOptions := &terraform.Options{

		TerraformDir: terraformDirectoryPath,
		NoColor:      true,
	}

	return terraformOptions
}
