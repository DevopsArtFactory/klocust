package kube

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	. "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeployment(namespace string, name string) (*v1.Deployment, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func GetDeployments(namespace string) (*v1.DeploymentList, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	deployments, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func IsDeploymentExists(namespace string, name string) (bool, error) {
	deployment, err := GetDeployment(namespace, name)
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
