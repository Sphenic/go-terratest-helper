package platformProvider

func GetPlatformProviderTestProjectCode(platformProviderCode string) string {

	platformProviderTestProjectCodes := map[string]string{
		//"AWS": "",
		//"GCP": "",
		//"MAZ": "",
	}

	return platformProviderTestProjectCodes[platformProviderCode]
}
