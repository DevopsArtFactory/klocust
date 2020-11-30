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
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	serializeryaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

var (
	handleFuncs = map[string]func(context.Context, kubernetes.Interface, *unstructured.Unstructured, []byte) error{
		"configmap":  configmap,
		"deployment": deployment,
		"service":    service,
		"ingress":    ingress,
	}
)

func configmap(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.CoreV1().ConfigMaps(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func deployment(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.AppsV1().Deployments(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func service(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.CoreV1().Services(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func ingress(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	if _, err := client.NetworkingV1beta1().Ingresses(obj.GetNamespace()).Patch(
		ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{FieldManager: "klocust"}); err != nil {
		return err
	}
	return nil
}

func Apply(namespace, filename string) (*unstructured.Unstructured, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var typeMeta runtime.TypeMeta
	if err := yaml.Unmarshal(bytes, &typeMeta); err != nil {
		return nil, fmt.Errorf("%v, decode yaml failed", err)
	}
	if typeMeta.Kind == "" {
		return nil, fmt.Errorf("%v, type kind is empty", err)
	}

	// Decode to unstructured object
	obj := &unstructured.Unstructured{}
	dec := serializeryaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	_, _, err = dec.Decode(bytes, nil, obj)
	if err != nil {
		return nil, fmt.Errorf("%v, unmarshal yaml failed", err)
	}

	f, ok := handleFuncs[strings.ToLower(obj.GetKind())]
	if !ok {
		return nil, fmt.Errorf("unsupported kind: %s", obj.GetKind())
	}

	if err := f(context.TODO(), client, obj, bytes); err != nil {
		return nil, err
	}

	return obj, nil
}
