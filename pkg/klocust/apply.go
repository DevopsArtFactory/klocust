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
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/DevopsArtFactory/klocust/pkg/kube"
	"github.com/DevopsArtFactory/klocust/pkg/kube/handler"
	"github.com/DevopsArtFactory/klocust/pkg/printer"
	"github.com/DevopsArtFactory/klocust/pkg/schemas"
)

func renderProjectTemplates(locustName string) ([]*bytes.Buffer, error) {
	var renderedBufList []*bytes.Buffer

	yamlFile, err := os.ReadFile(getLocustConfigFilename(locustName))
	if err != nil {
		return nil, err
	}

	var values schemas.LocustValues
	if err := yaml.Unmarshal(yamlFile, &values); err != nil {
		return nil, err
	}

	SetDefaultToValues(&values)

	for _, filename := range locustFilenames {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		renderedBuf, err := renderTemplateToBuf(
			getEmbedTemplatePath(filename),
			values)

		if err != nil {
			return nil, err
		}

		if renderedBuf == nil {
			continue
		}

		renderedBufList = append(renderedBufList, renderedBuf)
	}

	return renderedBufList, nil
}

func applyYamlFiles(out io.Writer, renderedBufList []*bytes.Buffer) error {
	for _, renderedBuf := range renderedBufList {
		obj, err := handler.Apply(renderedBuf)
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
	renderedBufList, err := renderProjectTemplates(locustName)
	if err != nil {
		return err
	}

	if err := applyYamlFiles(out, renderedBufList); err != nil {
		return err
	}

	printer.Default.Fprintf(out, "> End applying locust cluster: %s\n\n", locustName)

	if err := ListLocustDeployments(out, namespace, false); err != nil {
		return err
	}
	return nil
}

// SetDefaultToValues sets default values if user does not specify.
func SetDefaultToValues(values *schemas.LocustValues) {
	if values.Worker.Image == "" {
		values.Worker.Image = DefaultDockerImage
	}

	if values.Main.Image == "" {
		values.Main.Image = DefaultDockerImage
	}
}
