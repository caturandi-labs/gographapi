package console

import (
	"caturandi-labs/gographapi/command"
	"context"
	"log"
	"os"

	"github.com/erikgeiser/promptkit/textinput"
	"github.com/ggwhite/go-masker"
	"github.com/jedib0t/go-pretty/v6/table"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pterm/pterm"
)

func MembersOfGroupConsole(client *msgraphsdkgo.GraphServiceClient) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"No.", "ID", "Name", "Phones"})

	m := masker.New()

	input := textinput.New("Group ID : ")
	input.InitialValue = os.Getenv("AZURE_GROUP_ID")
	input.Placeholder = "Group ID is Required"
	userInput, err := input.RunPrompt()

	result, err := command.GetMembersByGroup(client, userInput)

	oDataCount := *result.GetOdataCount()

	pageIterator, err := msgraphgocore.NewPageIterator(result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)

	if err != nil {
		log.Fatalln("Error in Iterator ", err.Error())
	}

	var count, totalData int64

	count = 0
	totalData = 0
	loadMore := true

	err = pageIterator.Iterate(context.Background(), func(pageItem interface{}) bool {
		totalData++
		count++

		user := pageItem.(models.Userable)
		phones := ""
		if len(user.GetBusinessPhones()) > 0 {
			phones = user.GetBusinessPhones()[0]
		}
		t.AppendRows([]table.Row{
			{
				count,
				m.ID(*user.GetId()),
				m.Name(*user.GetDisplayName()),
				// m.Email(*user.GetUserPrincipalName()),
				phones,
			},
		})
		t.AppendSeparator()
		t.SetStyle(table.StyleLight)

		if (totalData%100 == 0) && (totalData < oDataCount) {
			t.Render()

			loadMore, err = pterm.DefaultInteractiveConfirm.
				WithDefaultText("Load More Data ?").
				Show()

			t.ResetRows()
		}

		return loadMore

	})

	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

}
