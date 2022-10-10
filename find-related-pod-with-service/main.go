package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", getEnv("KUBECONFIG", clientcmd.RecommendedHomeFile), "kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

    // Get nginx service Endpoints in default namespace
    // svc, err := clientset.CoreV1().Services("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
    endpoint, err := clientset.CoreV1().Endpoints("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
    fmt.Print(endpoint.Subsets)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
