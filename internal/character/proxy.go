package character

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/cache"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
)

type RickNMortyClientCacheProxy struct {
	ricknmorty.Client

	cache *cache.InMemCache[ricknmorty.PaginatedCharacters]
}

func NewRickNMortyClientCacheProxy(client ricknmorty.Client, cache *cache.InMemCache[ricknmorty.PaginatedCharacters]) ricknmorty.Client {
	return &RickNMortyClientCacheProxy{client, cache}
}

func (c *RickNMortyClientCacheProxy) FilterCharacters(ctx context.Context, filter ricknmorty.FilterCharactersQuery) (ricknmorty.PaginatedCharacters, error) {
	var (
		chars ricknmorty.PaginatedCharacters
		err   error
	)

	urlValues, err := query.Values(filter)
	if err != nil {
		return chars, err
	}

	key := urlValues.Encode()

	chars, err = c.cache.Get(ctx, key)
	found := err == nil
	if found {
		return chars, nil
	}

	chars, err = c.Client.FilterCharacters(ctx, filter)
	if err != nil {
		return chars, err
	}

	if err := c.cache.Set(ctx, key, chars); err != nil {
		return chars, err
	}

	return chars, nil
}
