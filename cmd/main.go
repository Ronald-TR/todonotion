package main

import (
	"fmt"

	"github.com/ronald-tr/todonotion"
)

func main() {
	client := todonotion.NewTodoClient()
	card := todonotion.NewCard("teste", "teste1")
	client.CreateCard(card, todonotion.LaneTypeToDo)
	fmt.Println(card.Page.ID)
	client.DeleteCardById(card.Page.ID)
}
