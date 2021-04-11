package terratest

import (
	"fmt"
	"testing"

	"github.com/sphenic/go-terratest-helper/terraform"
)

func RunTestPreActions(t *testing.T, testFilePath string, terraformVariables map[string]interface{}) {

	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("Test Script: %s\n", testFilePath)
	fmt.Printf("\n")

	// Save Terraform variables to file
	testVariablesFilePath := GetTestVariablesFilePath(terraform.GetTerraformRelativeDirectoryPath(), GetTestVariablesFileName())
	fmt.Printf("Saving Terraform variables to '%s'...\n", testFilePath)
	terraform.SaveTerraformVariablesToFile(testVariablesFilePath, terraformVariables)

}
