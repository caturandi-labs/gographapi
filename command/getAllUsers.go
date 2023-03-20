package command

import (
	"context"
	"log"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func GetAllUsers(client *msgraphsdkgo.GraphServiceClient) (models.UserCollectionResponseable, error) {
	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")

	requestCount := true
	var top int32
	top = 50

	options := &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Top:   &top,
			Count: &requestCount,
		},
	}

	result, err := client.Users().
		Get(context.Background(), options)
	if err != nil {
		log.Println("Error when query users ", err.Error())
		return nil, err
	}

	return result, nil
}
