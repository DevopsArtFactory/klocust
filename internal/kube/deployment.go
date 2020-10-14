package kube

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/clientcmd"
	"log"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetDeployments(namespace string) (*v1.DeploymentList, error) {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	deployments, err := client.AppsV1().Deployments(namespace).List(context.TODO(), meta_v1.ListOptions{})
	if err != nil {
		if errors.IsUnauthorized(err) {
			log.Fatal("Check your kubeconfig: ", err)
		}

		log.Fatal(err)
	}

	return deployments, nil
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
