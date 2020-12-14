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
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/DevopsArtFactory/klocust/pkg/kube"
	"github.com/DevopsArtFactory/klocust/pkg/kube/handler"
	"github.com/DevopsArtFactory/klocust/pkg/printer"
	"github.com/DevopsArtFactory/klocust/pkg/util"
)

func deleteFromYamlFiles(out io.Writer, namespace string, locustName string) error {
	for _, filename := range locustFilenames {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		filenameWithPath := getLocustProjectPath(locustName, filename)
		if !util.IsFileExists(filenameWithPath) {
			continue
		}

		obj, err := handler.Delete(namespace, filenameWithPath)
		if err != nil {
			return err
		}
		printer.Default.Fprintf(out, "%s `%s` deleted\n", strings.ToLower(obj.GetKind()), obj.GetName())
	}
	return nil
}

func DeleteLocust(out io.Writer, namespace string, locustName string) error {
	logrus.Debugf("Applied namespace: %s, Name: %s", namespace, locustName)
	if err := checkInitFileNotFound(locustName); err != nil {
		return err
	}

	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	mainDeploymentName := getLocustMainDeploymentName(locustName)
	if isExist, err := kube.IsDeploymentExists(namespace, mainDeploymentName); !isExist || err != nil {
		if !isExist {
			return fmt.Errorf("locust cluster not found: %s", locustName)
		}
		return nil
	}

	if err := deleteFromYamlFiles(out, namespace, locustName); err != nil {
		return err
	}

	return nil
}
