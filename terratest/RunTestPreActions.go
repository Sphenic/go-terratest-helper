package terratest

import (
	"fmt"
	"os"
	"testing"

	"github.com/sphenic/go-terratest-helper/terraform"
)

func RunTestPreActions(t *testing.T, testFilePath string, terraformVariables map[string]interface{}) {

	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	wd, _ := os.Getwd()
	fmt.Printf("WorkingDir: %s\n", wd)
	fmt.Printf("TestScript: %s\n", testFilePath)
	fmt.Printf("\n")

	// Save Terraform variables to file
	testVariablesFilePath := GetTestVariablesFilePath(terraform.GetTerraformRelativeDirectoryPath(), GetTestVariablesFileName())
	fmt.Printf("Saving Terraform variables to '%s'...\n", testVariablesFilePath)
	terraform.SaveTerraformVariablesToFile(testVariablesFilePath, terraformVariables)

}
