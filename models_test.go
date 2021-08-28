package todonotion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ronald-tr/todonotion"
)

var _ = Describe("Models", func() {
	var (
		card   *todonotion.Card
		client *todonotion.TodoClient
	)

	BeforeSuite(func() {
		card = todonotion.NewCard("TestCard", "test")
		client = todonotion.NewTodoClient()
	})
	Context("Card CRUD", func() {
		It("Create a new card in ToDo", func() {
			Expect(client.CreateCard(card, todonotion.LaneTypeToDo)).To(BeNil())
		})
		It("Move card to Doing", func() {
			Expect(client.MoveCard(card, todonotion.LaneTypeDoing)).To(BeNil())
		})
		It("Move card to Done", func() {
			Expect(client.MoveCard(card, todonotion.LaneTypeDone)).To(BeNil())
		})
		It("Remove card previously created", func() {
			Expect(client.DeleteCard(card)).To(BeNil())
		})
	})
})
