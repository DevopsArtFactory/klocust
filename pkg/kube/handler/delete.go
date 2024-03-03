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

package handler

import (
	"bytes"
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes"
)

var (
	deleteFuncs = ObjHandler{
		configmap:  deleteConfigmap,
		deployment: deleteDeployment,
		service:    deleteService,
		ingress:    deleteIngress,
	}
)

func deleteConfigmap(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, _ []byte) error {
	return client.CoreV1().ConfigMaps(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteDeployment(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, _ []byte) error {
	return client.AppsV1().Deployments(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteService(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, _ []byte) error {
	return client.CoreV1().Services(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteIngress(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, _ []byte) error {
	return client.NetworkingV1().Ingresses(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func Delete(renderedBuf *bytes.Buffer) (*unstructured.Unstructured, error) {
	return handleObjFromYamlFile(deleteFuncs, renderedBuf)
}
