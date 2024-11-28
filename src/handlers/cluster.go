package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetClusterDetails(c *gin.Context) {
	nodes, err := Clientset.CoreV1().Nodes().List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pods, err := Clientset.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clusterDetails := gin.H{
		"nodes": len(nodes.Items),
		"pods":  len(pods.Items),
		"status": c.GetHeader("status"),
		"IP": c.ClientIP(),
		"k8sVersion": c.GetHeader("k8s-version"),
		"cpu": c.GetHeader("cpu"),
		"memory": c.GetHeader("memory"),
		"storage": c.GetHeader("storage"),
	}

	c.JSON(http.StatusOK, clusterDetails)
}