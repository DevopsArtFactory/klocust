package klocust

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/schemas"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"k8s.io/klog/v2"
)

func checkInitFileNotFound(filenames []string, locustName string) error {
	for _, filename := range filenames {
		if isExist := util.IsFileExists(filename); !isExist {
			return errors.New(fmt.Sprintf(
				"`%s` file not found. \nYou need to init first before apply.\n\n$ klocust init %s\n",
				filename, locustName))
		}
	}
	return nil
}

func renderProjectTemplates(locustName, configFilename string) error {
	projectDir := getLocustProjectDir(locustName)

	if !util.IsDirExists(projectDir) {
		if err := util.CreateDir(projectDir); err != nil {
			return err
		}
	}
	yamlFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return err
	}

	var values schemas.LocustValues
	if err := yaml.Unmarshal(yamlFile, &values); err != nil {
		return err
	}

	// Create ./.klocust/{locustName}/{LocustFilenames}.yaml
	for _, filename := range locustFilenames {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		if _, err := renderTemplateFile(
			getLocustHomeTemplatesPath(filename),
			getLocustProjectPath(locustName, filename),
			values); err != nil {
			return err
		}
	}

	return nil
}

func applyYamlToK8s() {
}

func ApplyLocust(namespace string, locustName string) error {
	configFilename := getLocustConfigFilename(locustName)
	locustFilename := getLocustFilename(locustName)

	if err := checkInitFileNotFound([]string{configFilename, locustFilename}, locustName); err != nil {
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
		klog.Infof("> Start applying locust cluster: %s\n", locustName)
	} else {
		klog.Infof("> Start creating locust cluster: %s\n", locustName)
	}

	if err := renderProjectTemplates(locustName, configFilename); err != nil {
		return err
	}

	return nil
}
