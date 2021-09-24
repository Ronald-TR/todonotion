package entity

import (
	"github.com/dstotijn/go-notion"
)

type (
	LaneType   string
	FilterType string
)

const (
	LaneTypeToDo  LaneType = "To Do"
	LaneTypeDoing LaneType = "Doing"
	LaneTypeDone  LaneType = "Done"

	FilterTypeProperty FilterType = "Property"
	FilterTypeLane     FilterType = "Status"
	FilterTypeTitle    FilterType = "Name"
)

type Config struct {
	NotionKey  string
	DatabaseId string
}

type Card struct {
	Title    string
	Property string
	Page     notion.Page
	Lane     LaneType
}

func GetLaneType(lane string) LaneType {
	var output LaneType

	switch lane {
	case "To Do":
		output = LaneTypeToDo
	case "Doing":
		output = LaneTypeDoing
	case "Done":
		output = LaneTypeDone
	default:
		output = LaneTypeToDo
	}
	return output
}

func GetFilterType(lane string) LaneType {
	var output LaneType

	switch lane {
	case "To Do":
		output = LaneTypeToDo
	case "Doing":
		output = LaneTypeDoing
	case "Done":
		output = LaneTypeDone
	default:
		output = LaneTypeToDo
	}
	return output
}

func NewCard(title string, property string) *Card {
	return &Card{
		Title:    title,
		Property: property,
	}
}

// Set the LaneType by string, default "To Do"
func (c *Card) SetLane(lane string) {
	c.Lane = GetLaneType(lane)
}
