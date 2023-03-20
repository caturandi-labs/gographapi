package console

import (
	"caturandi-labs/gographapi/command"
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/textinput"
	"github.com/jedib0t/go-pretty/v6/table"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
)

func SingleUserConsole(client *msgraphsdkgo.GraphServiceClient) {
	input := textinput.New("ID/User Principal Name : ")
	input.InitialValue = ""
	input.Placeholder = "ID/User Principal Name is Required"
	userInput, err := input.RunPrompt()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	//"WengKen.Lee@taylors.edu.my"
	result, err := command.GetUserById(client, userInput)
	if err != nil {
		fmt.Printf("Error getting users: %v\n", err)
	}

	data := result

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Principal Name", "Phones"})
	t.AppendRows([]table.Row{
		{*data.GetId(), *data.GetDisplayName(), *data.GetUserPrincipalName(), *&data.GetBusinessPhones()[0]},
	})
	t.AppendSeparator()
	t.Render()
}
