package klocust

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/klog/v2"
)

func getLocustDeployments(namespace string) ([]v1.Deployment, error) {
	deployments, err := kube.GetDeployments(namespace)
	if err != nil {
		return nil, err
	}

	var locustDeployments = make([]v1.Deployment, 0)
	for _, deployment := range deployments.Items {
		if strings.HasPrefix(deployment.Name, locustMainDeploymentPrefix) {
			locustDeployments = append(locustDeployments, deployment)
		}
	}

	return locustDeployments, nil
}

func PrintLocustDeployments(namespace string) error {
	if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
		return err
	}

	locustDeployments, err := getLocustDeployments(namespace)
	if err != nil {
		return err
	}

	klog.Infof(">>> %d locust deployments in %s namespace. (PREFIX: %s)\n",
		len(locustDeployments), namespace, locustMainDeploymentPrefix)

	if len(locustDeployments) <= 0 {
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "DEPLOYMENT", "READY", "UP-TO-DATE", "AVAILABLE", "AGE"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, d := range locustDeployments {
		name := d.Name[len(locustMainDeploymentPrefix):]
		age := time.Since(d.CreationTimestamp.Time).Round(time.Second)

		table.Append([]string{
			name,
			d.Name,
			strconv.Itoa(int(d.Status.ReadyReplicas)) + "/" + strconv.Itoa(int(d.Status.Replicas)),
			strconv.Itoa(int(d.Status.UpdatedReplicas)),
			strconv.Itoa(int(d.Status.AvailableReplicas)),
			age.String()},
		)
	}
	table.Render()

	return nil
}
