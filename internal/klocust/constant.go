package klocust

import (
	"fmt"
	"os"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

type Locust struct {
	name             string
	namespace        string
	mainDeployment   appsv1.Deployment
	workerDeployment appsv1.Deployment
	ingress          v1beta1.Ingress
	service          v1.Service
}

var (
	userHomeDir, _                = os.UserHomeDir()
	locustHomeDir                 = userHomeDir + "/.klocust"
	locustHomeDefaultTemplatesDir = locustHomeDir + "/_default_templates"
)

const (
	locustlProjectDir                   = "./.klocust"
	locustProjectDefaultTemplatesDir    = locustlProjectDir + "/_default_templates"
	locustMainDeploymentPrefix          = "locust-main-"
	locustConfigFileSuffixWithExtension = "-klocust.yaml"
	locustFileSuffixWithExtension       = "-locustfile.py"

	locustFilename           = "locustfile.py"
	ingressFilename          = "main-ingress.yaml"
	serviceFilename          = "main-service.yaml"
	mainDeploymentFilename   = "main-deployment.yaml"
	workerDeploymentFilename = "worker-deployment.yaml"
	configMapFilename        = "configmap.yaml"
	valuesFilename           = "values.yaml"

	locustGitRepo = "https://raw.githubusercontent.com/DevopsArtFactory/klocust"

	DEFAULT_BUFFER_SIZE int64 = 1024
)

var locustFilenames = []string{
	locustFilename,

	valuesFilename,
	configMapFilename,
	mainDeploymentFilename,
	workerDeploymentFilename,
	serviceFilename,
	ingressFilename,
}

func getLocustGitRepoTemplatePath(filename string) string {
	subDir := "main/_default_templates/templates"
	if strings.HasSuffix(filename, ".py") {
		subDir = "main/_default_templates/tasks"
	}
	return fmt.Sprintf("%s/%s/%s", locustGitRepo, subDir, filename)
}

func getLocustHomeTemplatesPath(filename string) string {
	subDir := "templates"
	if strings.HasSuffix(filename, ".py") {
		subDir = "tasks"
	}
	return fmt.Sprintf("%s/%s/%s", locustHomeDefaultTemplatesDir, subDir, filename)
}

func getLocustProjectTemplatesPath(filename string) string {
	subDir := "templates"
	if strings.HasSuffix(filename, ".py") {
		subDir = "tasks"
	}
	return fmt.Sprintf("%s/%s/%s", locustProjectDefaultTemplatesDir, subDir, filename)
}

func getLocustProjectDir(locustName string) string {
	return fmt.Sprintf("%s/%s", locustlProjectDir, locustName)
}

func getLocustProjectPath(locustName string, filename string) string {
	return fmt.Sprintf("%s/%s", getLocustProjectDir(locustName), filename)
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
