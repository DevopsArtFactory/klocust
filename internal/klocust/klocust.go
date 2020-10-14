package klocust

import (
	"fmt"
	"github.com/DevopsArtFactory/klocust/internal/kube"
	v1 "k8s.io/api/apps/v1"
	"os"
	"strings"
	"text/tabwriter"
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

	fmt.Printf("%d locust deployments in %s namespace. (PREFIX: %s)\n",
		len(locustDeployments), namespace, LocustMasterDeploymentPrefix)

	if len(locustDeployments) <= 0 {
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	if _, err := fmt.Fprintf(writer, "\nDEPLOYMENT\tNAME\tREADY\tUP-TO-DATE\tAVAIABLE\tAGE\n"); err != nil {
		return err
	}

	for _, d := range locustDeployments {
		name := d.Name[len(LocustMasterDeploymentPrefix):]
		age := time.Since(d.CreationTimestamp.Time).Round(time.Second)
		if _, err := fmt.Fprintf(writer, "%s\t%s\t%d/%d\t%d\t%d\t%s\n",
			d.Name, name, d.Status.ReadyReplicas, d.Status.Replicas,
			d.Status.UpdatedReplicas, d.Status.AvailableReplicas, age); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
