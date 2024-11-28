package handlers

import (
    "github.com/gin-gonic/gin"
    "k8s.io/client-go/kubernetes"
)

var Clientset *kubernetes.Clientset

func SetClientset(clientset *kubernetes.Clientset) {
    Clientset = clientset
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// nodes apis routes
	r.GET("/api/nodes", GetNodes)
	r.POST("/api/nodes", CreateNode)
	r.GET("/api/nodes/:name", GetNode)
	r.DELETE("/api/nodes/:name", DeleteNode)
	r.PATCH("/api/nodes/:name", UpdateNode)

    // Pod  apis routes
    r.GET("/api/pods", ListPods)
    r.POST("/api/pods", CreatePod)
    r.GET("/api/pods/:name", GetPod)
    r.DELETE("/api/pods/:name", DeletePod)
    r.PATCH("/api/pods/:name", UpdatePod)

	// Cluster details
	r.GET("/api/cluster", GetClusterDetails)

	//deployment apis routes
	r.GET("/api/deployments", ListDeployments)
	r.POST("/api/deployments", CreateDeployment)
	r.GET("/api/deployments/:name", GetDeployment)
	r.DELETE("/api/deployments/:name", DeleteDeployment)
	r.PATCH("/api/deployments/:name", UpdateDeployment)
	


	return r
}