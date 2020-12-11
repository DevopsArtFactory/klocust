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
	"io"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/kube/handler"
	"github.com/DevopsArtFactory/klocust/internal/schemas"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"github.com/DevopsArtFactory/klocust/pkg/printer"
)

func renderProjectTemplates(locustName string) ([]string, error) {
	var renderedFileList []string

	projectDir := getLocustProjectDir(locustName)

	if !util.IsDirExists(projectDir) {
		if err := util.CreateDir(projectDir); err != nil {
			return nil, err
		}
	}
	yamlFile, err := ioutil.ReadFile(getLocustConfigFilename(locustName))
	if err != nil {
		return nil, err
	}

	var values schemas.LocustValues
	if err := yaml.Unmarshal(yamlFile, &values); err != nil {
		return nil, err
	}

	// Create ./.klocust/{locustName}/{...LocustFilenames}.yaml
	for _, filename := range locustFilenames {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		filePath := getLocustProjectPath(locustName, filename)

		renderedFile, err := renderTemplateFile(
			getLocustHomeTemplatesPath(filename),
			filePath,
			values)

		if err != nil {
			return nil, err
		}

		if renderedFile == "" {
			continue
		}

		renderedFileList = append(renderedFileList, renderedFile)
	}

	return renderedFileList, nil
}

func applyYamlFiles(out io.Writer, namespace string, yamlFiles []string) error {
	for _, filename := range yamlFiles {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		obj, err := handler.Apply(namespace, filename)
		if err != nil {
			return err
		}
		printer.Default.Fprintf(out, "%s `%s` configured\n", strings.ToLower(obj.GetKind()), obj.GetName())
	}
	return nil
}

// ApplyLocust creates a locust cluster with configuration files
func ApplyLocust(out io.Writer, namespace string, locustName string) error {
	logrus.Debugf("Applied namespace: %s, Name: %s", namespace, locustName)
	if err := checkInitFileNotFound(locustName); err != nil {
		return err
	}

	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	mainDeploymentName := getLocustMainDeploymentName(locustName)
	isExist, err := kube.IsDeploymentExists(namespace, mainDeploymentName)
	if err != nil {
		return err
	}

	if isExist {
		printer.Default.Fprintf(out, "> Start applying locust cluster: %s\n", locustName)
	} else {
		printer.Default.Fprintf(out, "> Start creating locust cluster: %s\n", locustName)
	}

	yamlFiles, err := renderProjectTemplates(locustName)
	if err != nil {
		return err
	}

	if err := applyYamlFiles(out, namespace, yamlFiles); err != nil {
		return err
	}

	printer.Default.Fprintf(out, "> End applying locust cluster: %s\n\n", locustName)

	if err := ListLocustDeployments(out, namespace, false); err != nil {
		return err
	}
	return nil
}
