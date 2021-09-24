package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	delCardCmd = &cobra.Command{
		Use:   "del",
		Short: "Delete one card from the Task board",
		Long:  "Given a entire or part of card title, finds the first ocurrence and then delete it.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			card.Title = args[0]
			card.SetLane(mvLane)
			err := client.FindCardByName(card)
			if err != nil {
				log.Fatalln(err)
			}
			err = client.DeleteCard(card)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v deleted.\n", card.Page.ID)
		},
	}
)

func init() {

}
