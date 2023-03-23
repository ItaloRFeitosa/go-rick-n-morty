package ricknmorty

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type FilterCharactersQuery struct {
	Name    string `url:"name" query:"name"`
	Status  string `url:"status" query:"status"`
	Species string `url:"species" query:"species"`
	Type    string `url:"type" query:"type"`
	Gender  string `url:"gender" query:"gender"`
	Page    int    `url:"page" query:"page"`
	Count   string `url:"count" query:"count"`
}

func (c *client) FilterCharacters(ctx context.Context, filter FilterCharactersQuery) (PaginatedCharacters, error) {
	var paginatedCharacters PaginatedCharacters

	urlValues, err := query.Values(filter)
	if err != nil {
		return paginatedCharacters, err
	}

	res, err := c.resty.R().
		SetContext(ctx).
		SetQueryParamsFromValues(urlValues).
		SetResult(PaginatedCharacters{}).
		Get("/character")

	if err != nil {
		return paginatedCharacters, fmt.Errorf("error on filter characters: %s: %w", err, ErrUnknown)
	}

	if res.IsError() {
		return paginatedCharacters, fmt.Errorf("error on filter characters; status code %d; body %s: %w", res.StatusCode(), res.String(), ErrResponse)
	}

	return *res.Result().(*PaginatedCharacters), nil
}
