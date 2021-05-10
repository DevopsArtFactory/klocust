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
	"io"

	"github.com/sirupsen/logrus"

	"github.com/DevopsArtFactory/klocust/pkg/kube"
	"github.com/DevopsArtFactory/klocust/pkg/printer"
	"github.com/DevopsArtFactory/klocust/pkg/schemas"
)

func createLocustProject(namespace string, locustName string) (string, string, error) {
	// Create ./{locustName}-klocust.yaml file
	var values schemas.LocustValues
	values.Namespace = namespace
	values.LocustName = locustName

	configFilename, err := renderValuesFile(
		getEmbedTemplatePath(valuesFilename),
		getLocustConfigFilename(locustName),
		values)
	if err != nil {
		return "", "", err
	}

	// Create ./{locustName}-locustfile.py file
	locustFilename, err := renderValuesFile(
		getEmbedTemplatePath(defaultLocustFilename),
		getLocustFilename(locustName),
		values)
	if err != nil {
		return "", "", err
	}

	return configFilename, locustFilename, nil
}

// InitLocust initialize locust files, not creating a cluster
func InitLocust(out io.Writer, namespace string, locustName string) error {
	logrus.Debugf("Applied namespace: %s, Name: %s", namespace, locustName)

	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	mainDeploymentName := getLocustMainDeploymentName(locustName)
	logrus.Debugf("Main deployment name generated: %s", mainDeploymentName)

	if isExist, err := kube.IsDeploymentExists(namespace, mainDeploymentName); isExist || err != nil {
		if isExist {
			return fmt.Errorf("`%s` deployment is already exists in `%s` namespace",
				mainDeploymentName, namespace)
		}
		return err
	}

	logrus.Debugf("Start to create locust project...")
	configFilename, locustFilename, err := createLocustProject(namespace, locustName)
	if err != nil {
		return err
	}

	printer.Green.Fprintf(out, "✓ %s has been successfully initialized!\n", locustName)
	printer.Default.Fprintf(out, "Please change `%s` and `%s` files.\n", configFilename, locustFilename)
	printer.Default.Fprintf(out, "And create locust cluster with next commands.\n\n")
	printer.Green.Fprintln(out, fmt.Sprintf("$ klocust apply %s", locustName))

	return nil
}
