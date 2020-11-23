package klocust

import (
	"errors"
	"github.com/DevopsArtFactory/klocust/internal/schemas"

	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/util"
)

func downloadDefaultTemplates() error {
	if err := util.CreateDir(locustHomeDefaultTemplatesDir); err != nil {
		return err
	}

	for _, filename := range locustFilenames {
		srcPath := getLocustGitRepoTemplatePath(filename)
		dstPath := getLocustHomeTemplatesPath(filename)

		if isExist := util.IsFileExists(dstPath); isExist {
			fmt.Printf("%s file exists already.\n", dstPath)
		}

		if err := util.DownloadFile(srcPath, dstPath); err != nil {
			return errors.New(fmt.Sprintf("Download Failed %s to %s: %v", srcPath, dstPath, err))
		}
		fmt.Printf("✓ %s file has downloaded.\n", dstPath)
	}

	return nil
}

func createLocustProject(namespace string, locustName string) (string, string, error) {
	projectDir := getLocustProjectDir(locustName)
	if err := util.CreateDir(projectDir); err != nil {
		return "", "", err
	}

	// Create ./{locustName}-values.yaml file
	var values schemas.LocustValues
	values.Namespace = namespace
	values.LocustName = locustName

	configFilename, err := renderValuesFile(
		getLocustHomeTemplatesPath(valuesFilename),
		getLocustConfigFilename(locustName),
		values)

	if err != nil {
		return "", "", err
	}

	// Create ./{locustName}-locustfile.py file
	locustFilename, err := util.CopyFile(
		getLocustHomeTemplatesPath(locustFilename),
		getLocustFilename(locustName),
		DEFAULT_BUFFER_SIZE)

	if err != nil {
		return "", "", err
	}

	return configFilename, locustFilename, nil
}

func InitLocust(namespace string, locustName string) error {
	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	masterDeploymentName := getLocustMasterDeploymentName(locustName)
	if isExist, err := kube.IsDeploymentExists(namespace, masterDeploymentName); isExist || err != nil {
		if isExist {
			return errors.New(fmt.Sprintf("`%s` deployment is already exists in `%s` namespace.",
				masterDeploymentName, namespace))
		}
		return err
	}

	if !util.IsDirExists(locustHomeDefaultTemplatesDir) {
		if err := downloadDefaultTemplates(); err != nil {
			return err
		}
	}

	configFilename, locustFilename, err := createLocustProject(namespace, locustName)
	if err != nil {
		return err
	}

	fmt.Printf("\n✓ %s has been successfully initialized!\n", locustName)
	fmt.Printf("Please change `%s` and `%s` files.\n", configFilename, locustFilename)
	fmt.Printf("And create locust cluster with next commands.\n\n")
	fmt.Printf("$ klocust apply %s -n %s\n", locustName, namespace)

	return nil
}
