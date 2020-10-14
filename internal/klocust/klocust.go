package klocust

import (
	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/api/apps/v1"
	"os"
	"strconv"
	"strings"
	"time"
)

const LocustMasterDeploymentPrefix = "locust-master-"

func getLocustDeployments(namespace string) ([]v1.Deployment, error) {
	deployments, err := kube.GetDeployments(namespace)
	if err != nil {
		return nil, err
	}

	var locustDeployments = make([]v1.Deployment, 0)
	for _, deployment := range deployments.Items {
		if strings.HasPrefix(deployment.Name, LocustMasterDeploymentPrefix) {
			locustDeployments = append(locustDeployments, deployment)
		}
	}

	return locustDeployments, nil
}

func PrintLocustDeployments(namespace string) error {
	if namespace == "" {
		var err error
		namespace, err = kube.GetNamespaceFromCurrentContext()
		if err != nil {
			return err
		}
	}

	locustDeployments, err := getLocustDeployments(namespace)
	if err != nil {
		return err
	}

	fmt.Printf(">>> %d locust deployments in %s namespace. (PREFIX: %s)\n",
		len(locustDeployments), namespace, LocustMasterDeploymentPrefix)

	if len(locustDeployments) <= 0 {
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"DEPLOYMENT", "NAME", "READY", "UP-TO-DATE", "AVAILABLE", "AGE"})
	table.SetHeaderAlignment(tablewriter.ALIGN_RIGHT)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	for _, d := range locustDeployments {
		name := d.Name[len(LocustMasterDeploymentPrefix):]
		age := time.Since(d.CreationTimestamp.Time).Round(time.Second)

		table.Append([]string{
			d.Name,
			name,
			strconv.Itoa(int(d.Status.ReadyReplicas)) + "/" + strconv.Itoa(int(d.Status.Replicas)),
			strconv.Itoa(int(d.Status.UpdatedReplicas)),
			strconv.Itoa(int(d.Status.AvailableReplicas)),
			age.String()},
		)
	}
	table.Render()

	return nil
}
