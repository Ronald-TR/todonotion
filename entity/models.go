package entity

import (
	"github.com/dstotijn/go-notion"
)

type Card struct {
	Title    string
	Property string
	Page     notion.Page
}

func NewCard(title string, property string) *Card {
	return &Card{
		Title:    title,
		Property: property,
	}
}
