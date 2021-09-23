package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	newCardCmd = &cobra.Command{
		Use:   "new",
		Short: "Create card in Task board",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			card.Title = args[0]
			card.SetLane(lane)
			fmt.Printf("%v created.\n", card.Title)
			client.CreateCard(card, card.Lane)
		},
	}
)

func init() {

	newCardCmd.PersistentFlags().StringVarP(&card.Property, "prop", "p", "todonotion", "Property Select value that you want to tag.")
	newCardCmd.PersistentFlags().StringVarP(&lane, "lane", "l", "To Do", "Lane to create the card.")
}
