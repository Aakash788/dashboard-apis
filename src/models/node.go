package models

import (
)

// Struct for creating a node
type CreateNodeRequest struct {
	Name string `json:"name" binding:"required"`
	Labels map[string]string `json:"labels"`
	IP string `json:"ip"`
	Role string `json:"role"`
	ProviderID string `json:"providerID"`
}

type UpdateNodeRequest struct {	
	Name string `json:"name" binding:"required"`
	Labels map[string]string `json:"labels"`
	IP string `json:"ip"`
	Role string `json:"role"`
	ProviderID string `json:"providerID"`
}