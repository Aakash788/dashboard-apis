package models

import "k8s.io/api/core/v1"

// Struct for creating a pod
type CreatePodRequest struct {
    Namespace  string            `json:"namespace" binding:"required"`
    Name       string            `json:"name" binding:"required"`
    Labels     map[string]string `json:"labels"`
    Containers []v1.Container    `json:"containers" binding:"required"`
	
}

// Struct for updating a pod

type UpdatePodRequest struct {
    Namespace string            `json:"namespace"`
	Name      string            `json:"name"`
	Labels    map[string]string `json:"labels"`
	Containers []v1.Container   `json:"containers"`
}
