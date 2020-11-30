package kube

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	serializeryaml "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

var (
	deleteFuncs = map[string]func(context.Context, kubernetes.Interface, *unstructured.Unstructured) error{
		"configmap":  deleteConfigmap,
		"deployment": deleteDeployment,
		"service":    deleteService,
		"ingress":    deleteIngress,
	}
)

func deleteConfigmap(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured) error {
	if err := client.CoreV1().ConfigMaps(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func deleteDeployment(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured) error {
	if err := client.AppsV1().Deployments(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func deleteService(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured) error {
	if err := client.CoreV1().Services(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func deleteIngress(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured) error {
	if err := client.NetworkingV1beta1().Ingresses(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func Delete(namespace, filename string) (*unstructured.Unstructured, error) {
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
		return nil, errors.New(fmt.Sprintf("%v, Decode yaml failed.", err))
	}
	if typeMeta.Kind == "" {
		return nil, errors.New(fmt.Sprintf("%v, Type kind is empty.", err))
	}

	// Decode to unstructured object
	obj := &unstructured.Unstructured{}
	dec := serializeryaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	_, _, err = dec.Decode(bytes, nil, obj)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v, Unmarshal yaml failed.", err))
	}

	f, ok := deleteFuncs[strings.ToLower(obj.GetKind())]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unsupported kind: %s.", obj.GetKind()))
	}

	if err := f(context.TODO(), client, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
