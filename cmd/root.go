package cmd

import (
	"fmt"

	entity "github.com/ronald-tr/todonotion/entity"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	client entity.TodoClient

	rootCmd = &cobra.Command{
		Use:   "todonotion",
		Short: "Todonotion is a todo replication inside Notion.so",
		Long: `A quick way to generate cards that needs to be done without leaving the terminal.
				Full documentation here: https://github.com/ronald-tr/todonotion
			`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			cmd.Help()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	cd := entity.NewCard("", "")
	client := entity.NewTodoClient()

	card := &cobra.Command{
		Use:   "todo",
		Short: "Create the card into ToDo collumn",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cd.Title = args[0]
			cd.Property = args[1]
			fmt.Printf("%v\n", cd.Title)
			client.CreateCard(cd, entity.LaneTypeToDo)
		},
	}
	rootCmd.AddCommand(card)
}
