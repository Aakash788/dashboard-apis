package models
import (
	"k8s.io/api/core/v1"
)

// Struct for creating a deployments
type CreateDeploymentRequest struct {
	Name          string            `json:"name" binding:"required"`
	Namespace    string            `json:"namespace" binding:"required"`
	Labels        map[string]string `json:"labels"`
	Replicas      int64            `json:"replicas"`
	Selector      map[string]string `json:"selector"`
	TemplateLabels map[string]string `json:"templateLabels"`
	Containers    []v1.Container    `json:"containers" binding:"required"`
	
}

type UpdateDeploymentRequest struct {	
	Name          string            `json:"name" binding:"required"`
	Namespace   string            `json:"namespace" binding:"required"`
	Labels        map[string]string `json:"labels"`
	Replicas      int64             `json:"replicas"`
	Selector      map[string]string `json:"selector"`
	TemplateLabels map[string]string `json:"templateLabels"`
	Containers    []v1.Container    `json:"containers" binding:"required"`
}
