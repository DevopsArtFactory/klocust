package kube

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetDeployment(namespace string, name string) (*v1.Deployment, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func GetDeployments(namespace string) (*v1.DeploymentList, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	deployments, err := client.AppsV1().Deployments(namespace).List(context.TODO(), meta_v1.ListOptions{})
	if err != nil {
		return nil, err
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
