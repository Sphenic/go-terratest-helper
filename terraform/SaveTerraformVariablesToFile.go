package terraform

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTerraformVariablesToFile(filePath string, variables map[string]interface{}) {

	os.Remove(filePath)

	jsonFile, errOpenFile := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if errOpenFile != nil {
		log.Fatal(errOpenFile)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	errEncode := encoder.Encode(variables)
	if errEncode != nil {
		log.Fatal(errEncode)
		return
	}
}
