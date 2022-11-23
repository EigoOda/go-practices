package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
    "github.com/tidwall/gjson"
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

    // Get nginx service endpoint in default namespace
    endpoint, err := clientset.CoreV1().Endpoints("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
    // Decode struct
    jsonData, err := json.Marshal(endpoint.Subsets)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get service endpoint IPs
    podIps := gjson.Get(string(jsonData), "#.addresses.#.ip")
    fmt.Println(podIps)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
