package character

import (
	"context"
	"log"

	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
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

	alertManager  RickNMortyAlertManager
	alertManager2 RickNMortyAlertManager
}

func NewRickNMortyClientAlertDecorator(
	client ricknmorty.Client,
	alertManager RickNMortyAlertManager,
	alertManager2 RickNMortyAlertManager,
) ricknmorty.Client {

	return &RickNMortyClientAlertDecorator{client, alertManager, alertManager2}
}

func (c *RickNMortyClientAlertDecorator) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	chars, err := c.Client.FilterCharacters(ctx, filter)

	if err != nil {
		go c.alertManager.SendAlert(ctx, err)
		go c.alertManager2.SendAlert(ctx, err)
		return chars, err
	}

	return chars, nil
}

type RickNMortyClientErrorDecorator struct {
	ricknmorty.Client
}

func NewRickNMortyClientErrorDecorator(client ricknmorty.Client) ricknmorty.Client {
	return &RickNMortyClientErrorDecorator{client}
}

func (c *RickNMortyClientErrorDecorator) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	var (
		chars ricknmorty.PaginatedCharacters
		err   error
	)

	if filter.Page < 0 {
		return chars, ricknmorty.ErrInvalidParams
	}

	chars, err = c.Client.FilterCharacters(ctx, filter)
	if err != nil {
		return chars, err
	}

	if len(chars.Results) == 0 {
		return chars, ricknmorty.ErrNotFound
	}

	return chars, nil
}
