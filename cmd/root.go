package cmd

import (
	entity "github.com/ronald-tr/todonotion/entity"
	"github.com/spf13/cobra"
)

var (
	client  = entity.NewTodoClient()
	card    = entity.NewCard("", "")
	lane    string
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

	rootCmd.AddCommand(newCardCmd)
	rootCmd.AddCommand(mvCardCmd)
	rootCmd.AddCommand(mvpCardCmd)
}
