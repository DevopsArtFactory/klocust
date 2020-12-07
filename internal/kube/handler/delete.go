package handler

import (
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

func deleteConfigmap(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	return client.CoreV1().ConfigMaps(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteDeployment(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	return client.AppsV1().Deployments(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteService(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	return client.CoreV1().Services(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func deleteIngress(ctx context.Context, client kubernetes.Interface, obj *unstructured.Unstructured, data []byte) error {
	return client.NetworkingV1beta1().Ingresses(obj.GetNamespace()).Delete(
		ctx, obj.GetName(), metav1.DeleteOptions{})
}

func Delete(namespace, filename string) (*unstructured.Unstructured, error) {
	return handleObjFromYamlFile(deleteFuncs, filename)
}
