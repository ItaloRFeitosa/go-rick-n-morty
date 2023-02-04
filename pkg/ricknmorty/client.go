package ricknmorty

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	FilterCharacters(context.Context, FilterCharactersQuery) (PaginatedCharacters, error)
}

type client struct {
	resty *resty.Client
}

func New(baseURL string) Client {
	return &client{
		resty: resty.New().SetBaseURL(baseURL),
	}
}
