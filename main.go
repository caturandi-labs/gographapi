package main

import (
	"caturandi-labs/gographapi/console"
	"caturandi-labs/gographapi/utils"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	err := godotenv.Load()
	isEnd := false

	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	// Init Credentials
	cred, error := utils.AzureCredential()

	if error != nil {
		log.Fatal("Error Azure Credential")
	}

	area, _ := pterm.DefaultArea.Start()
	str, _ := pterm.DefaultBigText.
		WithLetters(
			putils.LettersFromStringWithStyle("GO-", pterm.NewStyle(pterm.FgCyan)),
			putils.LettersFromStringWithStyle("MSGRAPH", pterm.NewStyle(pterm.FgLightMagenta))).Srender()

	str = pterm.DefaultCenter.Sprint(str)
	area.Update(str)
	area.Stop()

	header := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan))
	pterm.DefaultCenter.Println(header.Sprint("created by caturandi-labs"))

	// Client Adapter
	client, err := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})

	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	for !isEnd {

		options := []string{"Find User By ID/Principal Name", "Find All User", "Find Members by Group ID", "Quit Application"}
		result, _ := pterm.DefaultInteractiveSelect.
			WithOptions(options).
			WithDefaultText("Select Menu : ").
			Show()

		switch result {
		case options[0]:
			header := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgGreen))
			pterm.DefaultCenter.Println(header.Sprint("Find User By User Principal Name / Object ID"))

			console.SingleUserConsole(client)
		case options[1]:
			header := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgYellow))
			pterm.DefaultCenter.Println(header.Sprint("Find All Users"))

			console.AllUserConsole(client)
		case options[2]:

			header := pterm.DefaultHeader.WithFullWidth(true).WithBackgroundStyle(pterm.NewStyle(pterm.BgMagenta))
			pterm.DefaultCenter.Println(header.Sprint("Find Member by Group"))
			console.MembersOfGroupConsole(client)
		case options[3]:
			os.Exit(0)
		}

		exitApp, err := pterm.DefaultInteractiveConfirm.
			WithDefaultText("Exit Application ?").
			Show()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// Exit with Success (normal signal termination)
		if exitApp {
			os.Exit(0)
		}

		// Clear Screen (UNIX)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
