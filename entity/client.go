package entity

import (
	"context"
	"fmt"
	"os"

	"github.com/dstotijn/go-notion"
)

type TodoClient struct {
	client *notion.Client
	cfg    *Config
}

func NewConfig() *Config {
	return &Config{
		NotionKey:  os.Getenv("NOTION_KEY"),
		DatabaseId: os.Getenv("NOTION_DATABASE_ID"),
	}
}

func NewTodoClient() *TodoClient {
	cfg := NewConfig()
	return &TodoClient{
		client: notion.NewClient(cfg.NotionKey),
		cfg:    cfg,
	}
}
func (c *TodoClient) CreateCard(card *Card, lane LaneType) (err error) {
	page, err := c.client.CreatePage(context.Background(), notion.CreatePageParams{
		ParentType: notion.ParentTypeWorkspace,
		ParentID:   c.cfg.DatabaseId,
		DatabasePageProperties: &notion.DatabasePageProperties{
			"Status": {
				Select: &notion.SelectOptions{Name: string(lane)},
			},
			"Property": {
				Select: &notion.SelectOptions{Name: card.Property},
			},
			"Name": {Title: []notion.RichText{{Text: &notion.Text{Content: card.Title}}}},
		},
	})
	card.Page = page
	return err
}

func (c *TodoClient) DeleteCard(card *Card) error {
	page, err := c.client.UpdatePageProps(context.Background(), card.Page.ID, notion.UpdatePageParams{
		DatabasePageProperties: &notion.DatabasePageProperties{},
		Archived:               true,
	})
	card.Page = page
	return err
}

func (c *TodoClient) DeleteCardById(id string) error {
	_, err := c.client.UpdatePageProps(context.Background(), id, notion.UpdatePageParams{
		DatabasePageProperties: &notion.DatabasePageProperties{},
		Archived:               true,
	})
	return err
}

func (c *TodoClient) MoveCard(card *Card, to LaneType) error {
	_, err := c.client.UpdatePageProps(context.Background(), card.Page.ID, notion.UpdatePageParams{
		DatabasePageProperties: &notion.DatabasePageProperties{
			"Status": {
				Select: &notion.SelectOptions{Name: string(to)},
			},
		},
	})
	return err
}

// Injects inside the Card parameter the Page result
func (c *TodoClient) FindCardByName(card *Card) error {
	dbresult, err := c.client.QueryDatabase(context.Background(), c.cfg.DatabaseId, &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Name",
			Text:     &notion.TextDatabaseQueryFilter{Contains: card.Title}}},
	)
	if err != nil {
		return err
	}

	if len(dbresult.Results) == 0 {
		err = fmt.Errorf("no cards match for title that contains: %v", card.Title)
		return err
	}
	card.Page = dbresult.Results[0]

	return nil
}

// Injects inside the Card parameter the Page result
func (c *TodoClient) MapCardsByProperty(property string, fn func(card *Card)) error {
	dbresult, err := c.client.QueryDatabase(context.Background(), c.cfg.DatabaseId, &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Property",
			Select:   &notion.SelectDatabaseQueryFilter{Equals: property}}},
	)
	if err != nil {
		return err
	}

	for _, page := range dbresult.Results {
		props := page.Properties.(notion.DatabasePageProperties)
		card := NewCard(props["Name"].Title[0].PlainText, props["Property"].Select.Name)
		card.SetLane(props["Status"].Select.Name)
		card.Page = page
		fn(card)
	}

	return nil
}
