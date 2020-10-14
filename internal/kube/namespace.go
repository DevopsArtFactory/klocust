package kube

import (
	"k8s.io/client-go/tools/clientcmd"
)

func GetNamespaceFromCurrentContext() (string, error) {
	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return "", err
	}

	return clientCfg.Contexts[clientCfg.CurrentContext].Namespace, nil
}
