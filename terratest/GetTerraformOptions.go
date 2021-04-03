package terratest

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/sphenic/go-terratest-helper/terraformModule"
)

func GetTerraformOptions(variables map[string]interface{}) *terraform.Options {

	terraformOptions := &terraform.Options{

		TerraformDir: terraformModule.GetModuleRelativeDirectoryPath(),

		Vars: variables,

		NoColor: true,

	}

	return terraformOptions
}
