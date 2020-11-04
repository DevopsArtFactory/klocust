package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
