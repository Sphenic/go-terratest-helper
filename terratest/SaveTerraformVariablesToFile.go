package terratest

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTerraformVariablesToFile(variables map[string]interface{}, filePath string) {

	jsonFile, errOpenFile := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
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
