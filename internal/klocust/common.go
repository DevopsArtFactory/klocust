package klocust

func getLocustMasterDeploymentName(locustName string) string {
	return LocustMasterDeploymentPrefix + locustName
}

func getLocustFilename(locustName string) string {
	return locustName + LocustFileWithExtension
}

func getLocustConfigFilename(locustName string) string {
	return locustName + LocustConfigFileWithExtension
}
