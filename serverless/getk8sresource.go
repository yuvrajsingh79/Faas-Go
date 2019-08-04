package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	//_ "k8s.io/client-go/plugin/pkg/client/auth"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClientConfig first tries to get a config object which uses the service account kubernetes gives to pods,
// if it is called from a process running in a kubernetes environment.
// Otherwise, it tries to build config from a default kubeconfig filepath if it fails, it fallback to the default config.
// Once it get the config, it returns the same.
func GetClientConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Unable to create config. Error: %+v", err)
		err1 := err
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			err = fmt.Errorf("InClusterConfig as well as BuildConfigFromFlags Failed. Error in InClusterConfig: %+v\nError in BuildConfigFromFlags: %+v", err1, err)
			return nil, err
		}
	}

	return config, nil
}

// GetClientset first tries to get a config object which uses the service account kubernetes gives to pods,
// if it is called from a process running in a kubernetes environment.
// Otherwise, it tries to build config from a default kubeconfig filepath if it fails, it fallback to the default config.
// Once it get the config, it creates a new Clientset for the given config and returns the clientset.
func GetClientset() (*kubernetes.Clientset, error) {
	config, err := GetClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		err = fmt.Errorf("Failed creating clientset. Error: %+v", err)
		return nil, err
	}

	return clientset, nil
}

// PrettyString returns the prettified string of the interface supplied. (If it can)
func PrettyString(in interface{}) string {
	jsonStr, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		err := fmt.Errorf("Unable to marshal, Error: %+v", err)
		if err != nil {
			fmt.Printf("Unable to marshal, Error: %+v\n", err)
		}
		return fmt.Sprintf("%+v", in)
	}
	return string(jsonStr)
}

//Resources route handler function
func GetAllK8sresources(w http.ResponseWriter, r *http.Request) {
	clientset, err := GetClientset()
	if err != nil {
		panic(err)
	}

	param := r.URL.Path
	switch param {
	case "namespace":
		listNamespace(clientset)
	case "deployment":
		listDeployment(clientset)
	case "service":
		listService(clientset)
	case "pv":
		listPv(clientset)
	case "pvc":
		listPvc(clientset)
	case "pods":
		listPods(clientset)
	default:
		listNamespace(clientset)
		listDeployment(clientset)
		listService(clientset)
		listPv(clientset)
		listPvc(clientset)
		listPods(clientset)

	}

}

func listNamespace(clientset *kubernetes.Clientset) {
	namespaces, err := clientset.CoreV1().Namespaces().List(meta_v1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, n := range namespaces.Items {
		fmt.Println("Namespace: ", n.Name)
		fmt.Printf(PrettyString(n))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}

func listService(clientset *kubernetes.Clientset) {

	services, err := clientset.CoreV1().Services("").List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, s := range services.Items {
		fmt.Println("Service Name: ", s.Name)
		fmt.Printf(PrettyString(s))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}

func listPvc(clientset *kubernetes.Clientset) {

	pvcs, err := clientset.CoreV1().PersistentVolumeClaims("").List(meta_v1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, pvc := range pvcs.Items {
		fmt.Println("PVC Name: ", pvc.Name, "PVC StorageClass: ", pvc.Spec.StorageClassName)
		fmt.Printf(PrettyString(pvc))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}

func listDeployment(clientset *kubernetes.Clientset) {

	deployment, err := clientset.AppsV1().Deployments("").List(meta_v1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, d := range deployment.Items {
		fmt.Println("Deployment Name: ", d.Name)
		fmt.Printf(PrettyString(d))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}
func listPods(clientset *kubernetes.Clientset) {

	pods, err := clientset.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		fmt.Println("Pod Name: ", pod.Name)
		fmt.Printf(PrettyString(pod))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}

func listPv(clientset *kubernetes.Clientset) {

	pvs, err := clientset.CoreV1().PersistentVolumes().List(meta_v1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, pv := range pvs.Items {
		fmt.Println("PV Name: ", pv.Name, "PVC Name: ", pv.Spec.ClaimRef.Name, "PVC Namespace: ", pv.Spec.ClaimRef.Namespace)
		fmt.Printf(PrettyString(pv))
		fmt.Println()
		fmt.Println(strings.Repeat("*", 80))
	}
}
