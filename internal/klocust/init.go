package klocust

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"gopkg.in/yaml.v3"
	. "k8s.io/apimachinery/pkg/api/errors"
)

type locustConfig struct {
	Locust struct {
		Name   string `yaml:"name"`
		Master struct {
			Cpu    string `yaml:"cpu"`
			Memory string `yaml:"memory"`
		} `yaml:"master"`
		Worker struct {
			Count  int32  `yaml:"count"`
			Cpu    string `yaml:"cpu"`
			Memory string `yaml:"memory"`
		} `yaml:"worker"`
	}
}

func isExistLocustSet(namespace string, name string) (bool, error) {
	deployment, err := kube.GetDeployment(namespace, name)
	if err != nil {
		if e, ok := err.(*StatusError); ok && e.ErrStatus.Code == 404 {
			return false, nil
		}
		return false, err
	}

	if deployment != nil {
		return true, nil
	}

	return false, nil
}

func createConfigFile(kLocustName string) (filename string, err error) {
	filename = GetKLocustConfigFileName(kLocustName)

	if isExist := util.IsExistsFile(filename); isExist {
		return "", errors.New(fmt.Sprintf("`%s` file is already exists.", filename))
	}

	L := locustConfig{}
	L.Locust.Name = kLocustName
	L.Locust.Master.Cpu = LocustMasterDefaultCPU
	L.Locust.Master.Memory = LocustMasterDefaultMemory
	L.Locust.Worker.Count = LocustWorkerDefaultCount
	L.Locust.Worker.Cpu = LocustWorkerDefaultCPU
	L.Locust.Worker.Memory = LocustWorkerDefaultMemory

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(&L); err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	fmt.Printf("> %s file has created.\n", filename)

	return filename, nil
}

func createLocustFile(kLocustName string) (filename string, err error) {
	filename = GetLocustFileName(kLocustName)

	if isExist := util.IsExistsFile(filename); isExist {
		return "", errors.New(fmt.Sprintf("`%s` file is already exists.", filename))
	}

	if err := util.DownloadFile(DefaultLocustFileDownloadPath, filename); err != nil {
		return "", err
	}

	fmt.Printf("> %s file has created.\n", filename)

	return filename, nil
}

func InitLocust(namespace string, kLocustName string) error {
	if namespace == "" {
		var err error
		namespace, err = kube.GetNamespaceFromCurrentContext()
		if err != nil {
			return err
		}
	}

	isExist, err := isExistLocustSet(namespace, LocustMasterDeploymentPrefix+kLocustName)

	if err != nil {
		return err
	}

	if isExist {
		return errors.New(fmt.Sprintf("`%s%s` deployment is already exists in `%s` namespace.",
			LocustMasterDeploymentPrefix, kLocustName, namespace))
	}

	var (
		configFilename string
		locustFilename string
	)

	if configFilename, err = createConfigFile(kLocustName); err != nil {
		return err
	}

	if locustFilename, err = createLocustFile(kLocustName); err != nil {
		return err
	}

	fmt.Printf("\n%s has been successfully initialized!\n\n", kLocustName)
	fmt.Printf("Please change %s and %s files.\n", configFilename, locustFilename)
	fmt.Printf("And create locust cluster with next commands.\n\n")
	fmt.Printf("$ klocust create %s -n %s\n", kLocustName, namespace)

	return nil
}
