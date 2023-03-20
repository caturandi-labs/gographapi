package command

import (
	"context"
	"log"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func GetMembersByGroup(client *msgraphsdkgo.GraphServiceClient, groupId string) (models.DirectoryObjectCollectionResponseable, error) {

	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")

	requestCount := true
	var top int32
	top = 50

	result, err := client.GroupsById(groupId).Members().Get(context.Background(), &groups.ItemMembersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &groups.ItemMembersRequestBuilderGetQueryParameters{
			Count: &requestCount,
			Top:   &top,
		},
	})

	if err != nil {
		log.Println("Error when query members by group id : ", err.Error())
		return nil, err
	}

	return result, nil
}
