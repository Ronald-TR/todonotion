package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	mvLane    string
	mvCardCmd = &cobra.Command{
		Use:   "mv",
		Short: "Move the card in Task board",
		Long:  "Given a entire or part of card title, finds the first ocurrence and then move it.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			card.Title = args[0]
			card.SetLane(mvLane)
			err := client.FindCardByName(card)
			if err != nil {
				log.Fatalln(err)
			}
			err = client.MoveCard(card, card.Lane)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v moved.\n", card.Page.ID)
		},
	}
)

func init() {
	mvCardCmd.PersistentFlags().StringVarP(&mvLane, "lane", "l", "To Do", "Lane to move the card.")
}
