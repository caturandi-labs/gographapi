package utils

import (
	"fmt"
	"os"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)
func AzureCredential () (*azidentity.ClientSecretCredential, error) {
	opts := azidentity.ClientSecretCredentialOptions{}
	cred, err := azidentity.NewClientSecretCredential(os.Getenv("AZURE_TENANT_ID"), os.Getenv("AZURE_CLIENT_ID"), os.Getenv("AZURE_CLIENT_SECRET"), &opts)

	if err != nil {
		fmt.Printf("Error creating credentials: %v\n", err)
		return nil, err
	}

	return cred, nil
}