package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"azure-resource-api/models"
	"azure-resource-api/services"
)

type ResourceGroupHandler struct {
	azureService *services.AzureService
}

func NewResourceGroupHandler(azureService *services.AzureService) *ResourceGroupHandler {
	return &ResourceGroupHandler{
		azureService: azureService,
	}
}

// CreateResourceGroup godoc
// @Summary Create a new Azure Resource Group
// @Description Create a new resource group in the specified Azure subscription and location
// @Tags ResourceGroups
// @Accept json
// @Produce json
// @Param resourceGroup body models.ResourceGroupRequest true "Resource Group Details"
// @Success 201 {object} models.ResourceGroupResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resource-groups [post]
func (h *ResourceGroupHandler) CreateResourceGroup(c *gin.Context) {
	var req models.ResourceGroupRequest

	// Bind and validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create resource group
	resp, err := h.azureService.CreateResourceGroup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *ResourceGroupHandler) RegisterHandlers(router *gin.Engine) {
	router.POST("/resource-groups", h.CreateResourceGroup)
}
