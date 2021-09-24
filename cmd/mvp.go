package cmd

import (
	"fmt"

	entity "github.com/ronald-tr/todonotion/entity"
	"github.com/spf13/cobra"
)

var (
	from       string
	to         string
	mvpCardCmd = &cobra.Command{
		Use:   "mvp",
		Short: "Move all the cards in Task board between lanes",
		Long:  "Given a Property Select name, move all cards.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client.MapCardsBy(func(card *entity.Card) {
				lanefrom := entity.GetLaneType(from)
				laneto := entity.GetLaneType(to)
				if lanefrom == card.Lane {
					client.MoveCard(card, laneto)
					fmt.Printf("Card: %s moved to: %s\n", card.Title, laneto)
				}
			}, entity.CustomFilter{Type: entity.FilterTypeProperty, Value: args[0]})
		},
	}
)

func init() {
	mvpCardCmd.PersistentFlags().StringVarP(&from, "from", "f", "To Do", "Lane origin.")
	mvpCardCmd.PersistentFlags().StringVarP(&to, "to", "t", "Done", "Lane destination.")
}
