package command

import (
	"context"
	"log"

	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func GetUserById(client *msgraphsdkgo.GraphServiceClient, identity string) (models.Userable, error) {
	result, err := client.UsersById(identity).Get(context.Background(), nil)
	if err != nil {
		log.Println("Error when query user by id : ", err.Error())
		return nil, err
	}

	return result, nil
}
