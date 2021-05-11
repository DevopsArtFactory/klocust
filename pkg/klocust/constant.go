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
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/DevopsArtFactory/klocust/pkg/util"
)

var (
	locustFilenames = []string{
		defaultLocustFilename,

		valuesFilename,
		configMapFilename,
		mainDeploymentFilename,
		workerDeploymentFilename,
		serviceFilename,
		ingressFilename,
	}
)

const (
	locustMainDeploymentPrefix          = "locust-main-"
	locustConfigFileSuffixWithExtension = "-klocust.yaml"
	locustFileSuffixWithExtension       = "-locustfile.py"

	defaultLocustFilename    = "locustfile.py"
	ingressFilename          = "main-ingress.yaml"
	serviceFilename          = "main-service.yaml"
	mainDeploymentFilename   = "main-deployment.yaml"
	workerDeploymentFilename = "worker-deployment.yaml"
	configMapFilename        = "configmap.yaml"
	valuesFilename           = "values.yaml"

	// DefaultLogLevel is the default global verbosity
	DefaultLogLevel = logrus.WarnLevel
)

func getEmbedTemplatePath(filename string) string {
	subDir := "templates"
	if strings.HasSuffix(filename, ".py") {
		subDir = "tasks"
	}
	return fmt.Sprintf("_default_templates/%s/%s", subDir, filename)
}

func getLocustConfigFilename(locustName string) string {
	return locustName + locustConfigFileSuffixWithExtension
}
func getLocustFilename(locustName string) string {
	return locustName + locustFileSuffixWithExtension
}

func getLocustMainDeploymentName(locustName string) string {
	return locustMainDeploymentPrefix + locustName
}

// checkInitFileNotFound checks if there is init file
func checkInitFileNotFound(locustName string) error {
	filenames := []string{
		getLocustConfigFilename(locustName),
		getLocustFilename(locustName),
	}

	for _, filename := range filenames {
		if isExist := util.IsFileExists(filename); !isExist {
			return fmt.Errorf("`%s` file not found. \nYou need to init first before apply.\n\n$ klocust init %s", filename, locustName)
		}
	}
	return nil
}
