package entity_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	entity "github.com/ronald-tr/todonotion/entity"
)

var _ = Describe("entity", func() {
	var (
		card   *entity.Card
		client *entity.TodoClient
	)

	BeforeSuite(func() {
		card = entity.NewCard("TestCard", "test")
		client = entity.NewTodoClient()
	})
	Context("Card CRUD", func() {
		It("Create a new card in ToDo", func() {
			Expect(client.CreateCard(card, entity.LaneTypeToDo)).To(BeNil())
		})
		It("Move card to Doing", func() {
			Expect(client.MoveCard(card, entity.LaneTypeDoing)).To(BeNil())
		})
		It("Move card to Done", func() {
			Expect(client.MoveCard(card, entity.LaneTypeDone)).To(BeNil())
		})
		It("Remove card previously created", func() {
			Expect(client.DeleteCard(card)).To(BeNil())
		})
	})
})
