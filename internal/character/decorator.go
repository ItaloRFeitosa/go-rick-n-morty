package character

import (
	"context"
	"log"

	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
)

type RickNMortyClientLogProxy struct {
	ricknmorty.Client
}

func NewRickNMortyClientLogDecorator(client ricknmorty.Client) ricknmorty.Client {
	return &RickNMortyClientLogProxy{client}
}

func (c *RickNMortyClientLogProxy) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	chars, err := c.Client.FilterCharacters(ctx, filter)

	if err != nil {
		log.Println(err)
		return chars, err
	}

	return chars, nil
}
