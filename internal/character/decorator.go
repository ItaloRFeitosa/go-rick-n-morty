package character

import (
	"context"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
	"github.com/spf13/viper"
)

// Decorator Pattern Examples

type RickNMortyClientLogDecorator struct {
	ricknmorty.Client
}

func NewRickNMortyClientLogDecorator(client ricknmorty.Client) ricknmorty.Client {
	return &RickNMortyClientLogDecorator{client}
}

func (c *RickNMortyClientLogDecorator) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	chars, err := c.Client.FilterCharacters(ctx, filter)

	if err != nil {
		log.Println(err)
		return chars, err
	}

	return chars, nil
}

type RickNMortyClientAlertDecorator struct {
	ricknmorty.Client

	notificationClient *resty.Client
}

func NewRickNMortyClientAlertDecorator(client ricknmorty.Client) ricknmorty.Client {
	restyClient := resty.New()
	restyClient.SetBaseURL(viper.GetString("ALERT_URL"))
	return &RickNMortyClientAlertDecorator{client, restyClient}
}

func (c *RickNMortyClientAlertDecorator) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	chars, err := c.Client.FilterCharacters(ctx, filter)

	if err != nil {
		c.notificationClient.R().SetBody(map[string]string{"error": err.Error()}).Post("/")
		return chars, err
	}

	return chars, nil
}
