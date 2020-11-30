package klocust

import (
	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/kube"
	"k8s.io/klog/v2"
	"strings"
)

func deleteFromYamlFiles(namespace string, locustName string) error {
	for _, filename := range locustFilenames {
		if !strings.HasSuffix(filename, ".yaml") ||
			strings.HasSuffix(filename, valuesFilename) {
			continue
		}

		obj, err := kube.Delete(namespace, getLocustProjectPath(locustName, filename))
		if err != nil {
			return err
		}
		klog.Infof("%s `%s` deleted", strings.ToLower(obj.GetKind()), obj.GetName())
	}
	return nil
}

func DeleteLocust(namespace string, locustName string) error {
	configFilename := getLocustConfigFilename(locustName)
	locustFilename := getLocustFilename(locustName)

	if err := checkInitFileNotFound([]string{configFilename, locustFilename}, locustName); err != nil {
		return err
	}

	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	mainDeploymentName := getLocustMainDeploymentName(locustName)
	if isExist, err := kube.IsDeploymentExists(namespace, mainDeploymentName); !isExist || err != nil {
		if !isExist {
			return fmt.Errorf("%s locust cluster not found.", locustName)
		}
		return nil
	}

	if err := deleteFromYamlFiles(namespace, locustName); err != nil {
		return nil
	}

	return nil
}
