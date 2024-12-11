package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "azure-resource-api/docs" // Import generated docs
	"azure-resource-api/handlers"
	"azure-resource-api/services"
)

// @title Azure Resource Group Management API
// @version 1.0
// @description API for managing Azure Resource Groups
// @host localhost:8080
// @BasePath /
func main() {
	// Get subscription ID from environment variable
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		log.Fatal("AZURE_SUBSCRIPTION_ID must be set")
	}

	// Create Azure service
	azureService, err := services.NewAzureService(subscriptionID)
	if err != nil {
		log.Fatalf("Failed to create Azure service: %v", err)
	}

	// Create resource group handler
	resourceGroupHandler := handlers.NewResourceGroupHandler(azureService)

	// Setup Gin router
	router := gin.Default()

	// Register Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register handlers
	resourceGroupHandler.RegisterHandlers(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
