package cmd

import (
	"fmt"

	entity "github.com/ronald-tr/todonotion/entity"
	"github.com/spf13/cobra"
)

var (
	client  = entity.NewTodoClient()
	cd      = entity.NewCard("", "aaa")
	cardCmd = &cobra.Command{
		Use:   "new",
		Short: "Create the card in Task board",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cd.Title = args[0]
			fmt.Printf("%v created.\n", cd.Title)
			client.CreateCard(cd, entity.LaneTypeToDo)
		},
	}
)

func init() {

	cardCmd.PersistentFlags().StringVarP(&cd.Property, "prop", "p", "todonotion", "Property value that you want to use.")
}
