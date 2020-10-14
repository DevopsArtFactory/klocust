package kube

import (
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func GetNamespaceFromCurrentContext() string {
	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		log.Fatal(err)
	}

	return clientCfg.Contexts[clientCfg.CurrentContext].Namespace
}
