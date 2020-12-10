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
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	serializeryaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"

	"github.com/DevopsArtFactory/klocust/internal/kube"
)

var (
	configmap  = "configmap"
	deployment = "deployment"
	service    = "service"
	ingress    = "ingress"
)

type ObjHandler map[string]func(context.Context, kubernetes.Interface, *unstructured.Unstructured, []byte) error

func handleObjFromYamlFile(handler ObjHandler, filename string) (*unstructured.Unstructured, error) {
	client, err := kube.GetKubeClient()
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

	f, ok := handler[strings.ToLower(obj.GetKind())]

	if !ok {
		return nil, fmt.Errorf("unsupported kind: %s", obj.GetKind())
	}

	if err := f(context.TODO(), client, obj, bytes); err != nil {
		return nil, err
	}

	return obj, nil
}
