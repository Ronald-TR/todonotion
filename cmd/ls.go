package cmd

import (
	"fmt"

	entity "github.com/ronald-tr/todonotion/entity"
	"github.com/spf13/cobra"
)

var (
	lsLane     string
	lsTitle    string
	lsProperty string
	lsCardsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all cards in a given Lane from Task board",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s Cards:\n\n", lsLane)

			client.MapCardsBy(func(card *entity.Card) {
				fmt.Printf("%s @%s\n", card.Title, card.Property)
			},
				entity.CustomFilter{Type: entity.FilterTypeLane, Value: lsLane},
				entity.CustomFilter{Type: entity.FilterTypeProperty, Value: lsProperty},
				entity.CustomFilter{Type: entity.FilterTypeTitle, Value: lsTitle},
			)
		},
	}
)

func init() {
	lsCardsCmd.PersistentFlags().StringVarP(&lsLane, "lane", "l", "To Do", "Lane")
	lsCardsCmd.PersistentFlags().StringVarP(&lsTitle, "title", "t", "", "Card Title or part of it")
	lsCardsCmd.PersistentFlags().StringVarP(&lsProperty, "prop", "p", "", "Property Select tag.")
}
