package main

import (
    "log"
    "dashboard-api/src/handlers"
    "k8s.io/client-go/kubernetes"
)

var clientset *kubernetes.Clientset

func main() {
    var err error
    clientset, err = getK8sClient()
    if err != nil {
        log.Fatalf("Error creating Kubernetes client: %v", err)
    }

    handlers.SetClientset(clientset)

    r := handlers.SetupRouter()
    // Add more routes for pods and cluster details here

    log.Println("Starting server on :8080")
    r.Run(":8080")
}