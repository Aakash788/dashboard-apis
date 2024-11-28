package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/api/core/v1"
    "dashboard-api/src/models"
)

func ListPods(c *gin.Context) {
    namespace := c.DefaultQuery("namespace", "")
    pods, err := Clientset.CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list pods: " + err.Error()})
        return
    }

    podList := make([]map[string]string, 0)
    for _, pod := range pods.Items {
        podList = append(podList, map[string]string{
            "name":   pod.Name,
            "status": string(pod.Status.Phase),
        })
    }

    c.JSON(http.StatusOK, podList)
}

func CreatePod(c *gin.Context) {
    var req models.CreatePodRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
        return
    }

    pod := &v1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:   req.Name,
            Labels: req.Labels,
        },
        Spec: v1.PodSpec{
            Containers: req.Containers,
        },
    }

    createdPod, err := Clientset.CoreV1().Pods(req.Namespace).Create(c, pod, metav1.CreateOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pod: " + err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdPod)
}

func GetPod(c *gin.Context) {
    namespace := c.DefaultQuery("namespace", "")
    name := c.Param("name")
    pod, err := Clientset.CoreV1().Pods(namespace).Get(c, name, metav1.GetOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pod: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, pod)
}

func DeletePod(c *gin.Context) {
    namespace := c.DefaultQuery("namespace", "")
    name := c.Param("name")
    err := Clientset.CoreV1().Pods(namespace).Delete(c, name, metav1.DeleteOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pod: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pod deleted"})
}

func UpdatePod(c *gin.Context) {
    var req models.UpdatePodRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
        return
    }

    pod, err := Clientset.CoreV1().Pods(req.Namespace).Get(c, req.Name, metav1.GetOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pod: " + err.Error()})
        return
    }

    pod.Labels = req.Labels
    pod.Spec.Containers = req.Containers

    updatedPod, err := Clientset.CoreV1().Pods(req.Namespace).Update(c, pod, metav1.UpdateOptions{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pod: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedPod)
}