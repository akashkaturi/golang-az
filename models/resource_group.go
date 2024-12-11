package models

// ResourceGroupRequest represents the input for creating a resource group
type ResourceGroupRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

// ResourceGroupResponse represents the response after creating a resource group
type ResourceGroupResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	ID       string `json:"id"`
	Status   string `json:"status"`
}
