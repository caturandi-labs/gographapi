package console

import (
	"caturandi-labs/gographapi/command"
	"context"
	"log"
	"os"

	"github.com/ggwhite/go-masker"
	"github.com/jedib0t/go-pretty/v6/table"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pterm/pterm"
)

func AllUserConsole(client *msgraphsdkgo.GraphServiceClient) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetPageSize(2000)
	t.AppendHeader(table.Row{"ID", "Name", "Phones"})
	t.SetAutoIndex(true)

	m := masker.New()

	result, err := command.GetAllUsers(client)
	if err != nil {
		log.Fatalln("Error Get All users ", err.Error())
	}
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
