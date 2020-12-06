/*
Copyright 2020 The klocust Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package klocust

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/schemas"
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
			klog.Errorf("%s file exists already.\n", dstPath)
		}

		if err := util.DownloadFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("download Failed %s to %s: %v", srcPath, dstPath, err)
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
		getLocustHomeTemplatesPath(defaultLocustFilename),
		getLocustFilename(locustName),
		DefaultBufferSize)

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
			return fmt.Errorf("`%s` deployment is already exists in `%s` namespace",
				mainDeploymentName, namespace)
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
	klog.Infof("$ klocust apply %s", locustName)

	return nil
}
