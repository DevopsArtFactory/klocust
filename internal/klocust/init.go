package klocust

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/internal/util"
	"gopkg.in/yaml.v3"
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

func createConfigFile(locustName string) (filename string, err error) {
	filename = getLocustConfigFilename(locustName)

	if isExist := util.IsFileExists(filename); isExist {
		return "", NewFileExistsError(filename)
	}

	L := locustConfig{}
	L.Locust.Name = locustName
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

	if err = ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return "", err
	}

	fmt.Printf("> %s file has created.\n", filename)

	return filename, nil
}

func createLocustFile(locustName string) (filename string, err error) {
	filename = getLocustFilename(locustName)

	if isExist := util.IsFileExists(filename); isExist {
		return "", NewFileExistsError(filename)
	}

	if err := util.DownloadFile(DefaultLocustFileDownloadPath, filename); err != nil {
		return "", err
	}

	fmt.Printf("> %s file has created.\n", filename)

	return filename, nil
}

func InitLocust(namespace string, locustName string) error {
	var (
		masterDeploymentName string
		configFilename string
		locustFilename string
		err error
	)

	if namespace == "" {
		namespace, err = kube.GetNamespaceFromCurrentContext()
		if err != nil {
			return err
		}
	}

	masterDeploymentName = getLocustMasterDeploymentName(locustName)
	isExist, err := kube.IsDeploymentExists(namespace, masterDeploymentName)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New(fmt.Sprintf("`%s` deployment is already exists in `%s` namespace.",
			masterDeploymentName, namespace))
	}

	if configFilename, err = createConfigFile(locustName); err != nil {
		return err
	}

	if locustFilename, err = createLocustFile(locustName); err != nil {
		return err
	}

	fmt.Printf("\n%s has been successfully initialized!\n\n", locustName)
	fmt.Printf("Please change %s and %s files.\n", configFilename, locustFilename)
	fmt.Printf("And create locust cluster with next commands.\n\n")
	fmt.Printf("$ klocust create %s -n %s\n", locustName, namespace)

	return nil
}
