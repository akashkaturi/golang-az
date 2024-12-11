package services

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"azure-resource-api/models"
)

type AzureService struct {
	subscriptionID string
	client         *armresources.ResourceGroupsClient
}

func NewAzureService(subscriptionID string) (*AzureService, error) {
	// Create a credential using environment variables or Azure CLI
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create credential: %v", err)
	}

	// Create resource group client
	client, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource group client: %v", err)
	}

	return &AzureService{
		subscriptionID: subscriptionID,
		client:         client,
	}, nil
}

func (s *AzureService) CreateResourceGroup(ctx context.Context, req models.ResourceGroupRequest) (*models.ResourceGroupResponse, error) {
	// Prepare resource group parameters
	params := armresources.ResourceGroup{
		Location: to.Ptr(req.Location),
		Tags: map[string]*string{
			"ManagedBy": to.Ptr("GoAPI"),
		},
	}

	// Create the resource group
	resp, err := s.client.CreateOrUpdate(ctx, req.Name, params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource group: %v", err)
	}

	return &models.ResourceGroupResponse{
		Name:     *resp.ResourceGroup.Name,
		Location: *resp.ResourceGroup.Location,
		ID:       *resp.ResourceGroup.ID,
		Status:   "Created",
	}, nil
}
