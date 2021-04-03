package platformProviderGcp

import "strings"

func getGcpResourceLinkInRegion(gcpProjectProviderProjectCode string, gcpRegionCode string, gcpResourceTypePlural string, gcpResourceName string) string {

	return "projects/" +
	       strings.ToLower(gcpProjectProviderProjectCode) +"/"+
		   "regions/" +
		   strings.ToLower(gcpRegionCode) + "/" +
		   strings.ToLower(gcpResourceTypePlural) + "/" +
		   strings.ToLower(gcpResourceName)
}
