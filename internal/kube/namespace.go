package kube

import (
	"k8s.io/client-go/tools/clientcmd"
)

func GetNamespaceFromCurrentContext() (string, error) {
	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return "", err
	}

	if namespace := clientCfg.Contexts[clientCfg.CurrentContext].Namespace; namespace != "" {
		return namespace, nil
	}

	return "default", nil
}
