package entity

import "github.com/dstotijn/go-notion"

type CustomFilter struct {
	Type  FilterType
	Value string
}

func (f *CustomFilter) GetFilter() notion.DatabaseQueryFilter {
	var result notion.DatabaseQueryFilter
	switch f.Type {
	case FilterTypeLane:
		result = NewFilterLane(f.Value)
	case FilterTypeProperty:
		result = NewFilterProperty(f.Value)
	case FilterTypeTitle:
		result = NewFilterTitle(f.Value)
	}
	return result
}

func NewFilterProperty(value string) notion.DatabaseQueryFilter {
	return notion.DatabaseQueryFilter{
		Property: "Property",
		Select:   &notion.SelectDatabaseQueryFilter{Equals: value}}
}

func NewFilterLane(value string) notion.DatabaseQueryFilter {
	return notion.DatabaseQueryFilter{
		Property: "Status",
		Select:   &notion.SelectDatabaseQueryFilter{Equals: value}}
}

func NewFilterTitle(value string) notion.DatabaseQueryFilter {
	return notion.DatabaseQueryFilter{
		Property: "Name",
		Text:     &notion.TextDatabaseQueryFilter{Contains: value}}
}
