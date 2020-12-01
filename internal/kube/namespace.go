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

func SetCurrentNamespaceIfBlank(namespace *string) (string, error) {
	if *namespace == "" {
		var err error
		if *namespace, err = GetNamespaceFromCurrentContext(); err != nil {
			return "", err
		}
	}
	return *namespace, nil
}
