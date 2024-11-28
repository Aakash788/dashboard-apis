package handlers

import (
	"dashboard-api/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodes(c *gin.Context) {
    nodes, err := Clientset.CoreV1().Nodes().List(c, metav1.ListOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    nodeList := make([]map[string]string, 0)
    for _, node := range nodes.Items {
        nodeList = append(nodeList, map[string]string{
            "name":   node.Name,
            "status": string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
        })
    }

    c.JSON(http.StatusOK, nodeList)
}

func GetNode(c *gin.Context) {
	name := c.Param("name")
	node, err := Clientset.CoreV1().Nodes().Get(c, name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, node)
}

func CreateNode(c *gin.Context) {
	var req models.CreateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   req.Name,
			Labels: req.Labels,
		},
		Spec: v1.NodeSpec{
			ProviderID: req.ProviderID,
		},
	}

	createdNode, err := Clientset.CoreV1().Nodes().Create(c, node, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create node: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNode)
}

func DeleteNode(c *gin.Context) {
	name := c.Param("name")
	err := Clientset.CoreV1().Nodes().Delete(c, name, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
}

func UpdateNode(c *gin.Context) {
	name := c.Param("name")
	var req models.UpdateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	node, err := Clientset.CoreV1().Nodes().Get(c, name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	node.Labels = req.Labels
	updatedNode, err := Clientset.CoreV1().Nodes().Update(c, node, metav1.UpdateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedNode)
}
