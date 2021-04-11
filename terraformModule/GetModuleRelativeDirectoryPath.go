package terraformModule

import (
	"github.com/sphenic/go-terratest-helper/terraform"
)

func GetModuleRelativeDirectoryPath() string {
	return terraform.GetTerraformRelativeDirectoryPath()
}
