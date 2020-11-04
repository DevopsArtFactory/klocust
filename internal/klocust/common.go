package klocust

func getKLocustMasterDeploymentName(kLocustName string) string {
	return LocustMasterDeploymentPrefix + kLocustName
}

func getLocustFilename(kLocustName string) string {
	return kLocustName + LocustFileWithExtension
}

func getKLocustConfigFilename(kLocustName string) string {
	return kLocustName + LocustConfigFileWithExtension
}
