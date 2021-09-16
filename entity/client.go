package entity

import (
	"context"
	"os"

	"github.com/dstotijn/go-notion"
)

type (
	LaneType string
)

const (
	LaneTypeToDo  LaneType = "To Do"
	LaneTypeDoing LaneType = "Doing"
	LaneTypeDone  LaneType = "Done"
)

type Config struct {
	NotionKey  string
	DatabaseId string
}

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
				RichText: []notion.RichText{{Text: &notion.Text{Content: card.Property}}},
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
