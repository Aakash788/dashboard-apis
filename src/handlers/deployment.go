package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/apps/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"dashboard-api/src/models"
)

func getDeploymentStatus(conditions []v1.DeploymentCondition) string {
	for _, condition := range conditions {
		if condition.Type == v1.DeploymentAvailable {
			return string(condition.Status)
		}
	}
	return "Unknown"
}

func ListDeployments(c *gin.Context) {
	deployments, err := Clientset.AppsV1().Deployments("").List(c, metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deploymentList := make([]map[string]string, 0)
	for _, deployment := range deployments.Items {
		deploymentList = append(deploymentList, map[string]string{
			"name":   deployment.Name,
			"status": getDeploymentStatus(deployment.Status.Conditions),
		})
	}

	c.JSON(http.StatusOK, deploymentList)
}

func GetDeployment(c *gin.Context) {
	name := c.Param("name")
	deployment, err := Clientset.AppsV1().Deployments("").Get(c, name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deployment)
}

func int32Ptr(i int64) *int32 {
    v := int32(i)
    return &v
}

func CreateDeployment(c *gin.Context) {
	var req models.CreateDeploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   req.Name,
			Labels: req.Labels,
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(req.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: req.Selector,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: req.TemplateLabels,
				},
				Spec: corev1.PodSpec{
					Containers: req.Containers,
				},
			},
		},
	}

	createdDeployment, err := Clientset.AppsV1().Deployments(req.Namespace).Create(c, deployment, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deployment: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDeployment)
}


func DeleteDeployment(c *gin.Context) {
	name := c.Param("name")
	err := Clientset.AppsV1().Deployments("").Delete(c, name, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployment deleted"})
}

func UpdateDeployment(c *gin.Context) {
	name := c.Param("name")
	var req models.UpdateDeploymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	deployment, err := Clientset.AppsV1().Deployments("").Get(c, name, metav1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deployment.ObjectMeta.Labels = req.Labels
	deployment.Spec.Replicas = int32Ptr(req.Replicas)
	deployment.Spec.Selector.MatchLabels = req.Selector
	deployment.Spec.Template.ObjectMeta.Labels = req.TemplateLabels
	deployment.Spec.Template.Spec.Containers = req.Containers

	updatedDeployment, err := Clientset.AppsV1().Deployments("").Update(c, deployment, metav1.UpdateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update deployment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDeployment)
}