package klocust

import (
	"errors"
	"fmt"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/schemas"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"k8s.io/klog/v2"
)

func downloadDefaultTemplates() error {
	if err := util.CreateDir(locustHomeDefaultTemplatesDir); err != nil {
		return err
	}

	for _, filename := range locustFilenames {
		srcPath := getLocustGitRepoTemplatePath(filename)
		dstPath := getLocustHomeTemplatesPath(filename)

		if isExist := util.IsFileExists(dstPath); isExist {
			klog.Errorf("%s file exists already.\n", dstPath)
		}

		if err := util.DownloadFile(srcPath, dstPath); err != nil {
			return errors.New(fmt.Sprintf("Download Failed %s to %s: %v", srcPath, dstPath, err))
		}
		klog.Infof("✓ %s file has downloaded.\n", dstPath)
	}

	return nil
}

func createLocustProject(namespace string, locustName string) (string, string, error) {
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

	mainDeploymentName := getLocustMainDeploymentName(locustName)
	if isExist, err := kube.IsDeploymentExists(namespace, mainDeploymentName); isExist || err != nil {
		if isExist {
			return errors.New(fmt.Sprintf("`%s` deployment is already exists in `%s` namespace.",
				mainDeploymentName, namespace))
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

	klog.Infof("\n✓ %s has been successfully initialized!\n", locustName)
	klog.Infof("Please change `%s` and `%s` files.\n", configFilename, locustFilename)
	klog.Infof("And create locust cluster with next commands.\n\n")
	klog.Infof("$ klocust apply %s -n %s\n", locustName, namespace)

	return nil
}
