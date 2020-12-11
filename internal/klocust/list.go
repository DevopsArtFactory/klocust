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
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/api/apps/v1"

	"github.com/DevopsArtFactory/klocust/internal/kube"
	"github.com/DevopsArtFactory/klocust/pkg/printer"
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

// ListLocustDeployments shows list of locust cluster in kubernetes cluster
func ListLocustDeployments(out io.Writer, namespace string, allNamespace bool) error {
	if !allNamespace {
		if _, err := kube.SetCurrentNamespaceIfBlank(&namespace); err != nil {
			return err
		}
	} else {
		namespace = ""
	}

	locustDeployments, err := getLocustDeployments(namespace)
	if err != nil {
		return err
	}

	if allNamespace {
		namespace = "all"
	}

	printer.Green.Fprintf(out, ">>> %d locust deployments in %s namespace(s). (PREFIX: %s)\n",
		len(locustDeployments), namespace, locustMainDeploymentPrefix)

	if len(locustDeployments) == 0 {
		printer.Default.Fprintf(out, "No cluster exists in %s namespace(s)\n", namespace)
		printer.Default.Fprintln(out, "You can find clusters in other Namespaces with -A options")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(createListTableHeader(allNamespace))
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, d := range locustDeployments {
		name := d.Name[len(locustMainDeploymentPrefix):]
		age := time.Since(d.CreationTimestamp.Time).Round(time.Second)

		var values []string
		if allNamespace {
			values = append(values, d.Namespace)
		}

		table.Append(append(values, []string{
			name,
			d.Name,
			strconv.Itoa(int(d.Status.ReadyReplicas)) + "/" + strconv.Itoa(int(d.Status.Replicas)),
			strconv.Itoa(int(d.Status.UpdatedReplicas)),
			strconv.Itoa(int(d.Status.AvailableReplicas)),
			age.String()}...,
		))
	}
	table.Render()

	return nil
}

func createListTableHeader(allNamespace bool) []string {
	var header []string
	if allNamespace {
		header = append(header, "NAMESPACE")
	}
	return append(header, []string{"NAME", "DEPLOYMENT", "READY", "UP-TO-DATE", "AVAILABLE", "AGE"}...)
}
