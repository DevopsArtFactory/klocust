package klocust

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	"k8s.io/api/extensions/v1beta1"
)

const LocustMasterDeploymentPrefix = "locust-master-"

const LocustMasterDefaultCPU = "250m"
const LocustMasterDefaultMemory = "512Mi"

const LocustWorkerDefaultCount = 1
const LocustWorkerDefaultCPU = "250m"
const LocustWorkerDefaultMemory = "512Mi"

const LocustConfigFileWithExtension = "-klocust.yaml"
const LocustFileWithExtension = "-locustfile.py"

type Locust struct {
	name             string
	namespace        string
	masterDeployment appsv1.Deployment
	workerDeployment appsv1.Deployment
	ingress          v1beta1.Ingress
	service          v1.Service
}

const DefaultLocustFileDownloadPath = "https://raw.githubusercontent.com/DevopsArtFactory/klocust/main/examples/default.locustfile.py"
