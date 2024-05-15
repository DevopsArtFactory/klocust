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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

var (
	applyFuncs = ObjHandler{
		configmap:  applyConfigmap,
		deployment: applyDeployment,
		service:    applyService,
		ingress:    applyIngress,
	}
)

func applyConfigmap(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.CoreV1().ConfigMaps(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func applyDeployment(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.AppsV1().Deployments(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func applyService(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.CoreV1().Services(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func applyIngress(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.NetworkingV1().Ingresses(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func Apply(renderedBuf *bytes.Buffer) (*unstructured.Unstructured, error) {
	return handleObjFromYamlFile(applyFuncs, renderedBuf)
}
