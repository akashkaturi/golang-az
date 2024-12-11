# Azure Resource Group Management API

## Prerequisites
- Go 1.21+
- Azure CLI
- Azure Account

## Setup Instructions
1. Clone the repository
2. Install dependencies
3. Set up Azure credentials
4. Run the application

## Environment Variables
- `AZURE_SUBSCRIPTION_ID`: Your Azure Subscription ID
- `AZURE_TENANT_ID`: Azure Tenant ID
- `AZURE_CLIENT_ID`: Service Principal Client ID
- `AZURE_CLIENT_SECRET`: Service Principal Client Secret

## Running the Application
```bash
# Install dependencies
go mod tidy

# Generate Swagger docs
swag init

# Run the application
go run main.go


# Azure Resource Group Management API

## 📋 Overview

This Go-based API provides a streamlined interface for managing Azure Resource Groups, offering simple and efficient resource group creation through a RESTful endpoint. Built with Go, Gin Framework, and Azure SDK, the project includes Swagger documentation for easy integration and testing.

## ✨ Features

- 🚀 Create Azure Resource Groups programmatically
- 🔒 Secure Azure authentication
- 📖 Swagger UI documentation
- 🐳 Docker support
- 🧪 Structured error handling

## 🛠 Prerequisites

### Software Requirements
- Go 1.21+
- Azure CLI
- Docker (optional)
- Git

### Azure Requirements
- Active Azure Subscription
- Service Principal or Azure CLI access

## 🔧 Installation

### 1. Clone the Repository
```bash
git clone https://github.com/YOUR-USERNAME/azure-resource-group-api.git
cd azure-resource-group-api
```

### 2. Install Dependencies
```bash
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
swag init
```

## 🔐 Authentication Methods

### Option 1: Azure CLI Authentication (Development)
```bash
# Login to Azure
az login

# Set active subscription
az account set --subscription "YOUR-SUBSCRIPTION-ID"
```

### Option 2: Service Principal Authentication (Production)
```bash
# Create Service Principal
az ad sp create-for-rbac --name "APIServicePrincipal" --role Contributor
```

## 🚀 Running the Application

### Local Development
```bash
# Set environment variables
export AZURE_SUBSCRIPTION_ID=your-subscription-id
export AZURE_TENANT_ID=your-tenant-id
export AZURE_CLIENT_ID=your-client-id
export AZURE_CLIENT_SECRET=your-client-secret

# Run the application
go run main.go
```

### Docker Deployment
```bash
# Build Docker image
docker build -t azure-resource-api .

# Run Docker container
docker run -e AZURE_SUBSCRIPTION_ID=your-subscription-id \
           -e AZURE_TENANT_ID=your-tenant-id \
           -e AZURE_CLIENT_ID=your-client-id \
           -e AZURE_CLIENT_SECRET=your-client-secret \
           -p 8080:8080 \
           azure-resource-api
```

## 📡 Endpoints

### Create Resource Group
- **URL**: `/resource-groups`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "my-resource-group",
    "location": "eastus"
  }
  ```

## 📊 Swagger Documentation
Access Swagger UI at: `http://localhost:8080/swagger/index.html`

## 🧪 Testing

### Run Tests
```bash
go test ./...
```

## 🔒 Security Considerations
- Never commit sensitive credentials to version control
- Use environment variables or secure secret management
- Implement proper RBAC in Azure

## 🐛 Troubleshooting

### Common Issues
- Ensure Azure CLI is authenticated
- Verify subscription permissions
- Check environment variable configuration

## 📝 Contributing
1. Fork the repository
2. Create your feature branch
3. Commit changes
4. Push to the branch
5. Create a Pull Request

## 📄 License
This project is licensed under the Apache License 2.0

## 🤝 Support
For issues or questions, please open a GitHub issue in the repository.