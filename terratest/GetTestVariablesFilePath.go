package terratest

func GetTestVariablesFilePath(terraformDirectoryPath string, testVariablesFileName string) string {
	return terraformDirectoryPath + "/" + testVariablesFileName
}
