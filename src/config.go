package main

import (
    "flag"
    "fmt"
    "os"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func getK8sClient() (*kubernetes.Clientset, error) {
    
    kubeconfigPath := "../config/kubeconfig"
    if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
        return nil, fmt.Errorf("kubeconfig file does not exist at path: %s", kubeconfigPath)
    }

    kubeconfig := flag.String("kubeconfig", kubeconfigPath, "(optional) absolute path to the kubeconfig file")
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        return nil, err
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }

    return clientset, nil
}