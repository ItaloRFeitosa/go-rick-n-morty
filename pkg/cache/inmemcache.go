package cache

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Adapter Pattern Example

type InMemCache[T any] struct {
	gocache *gocache.Cache
}

func NewInMemCache[T any](defaultExp time.Duration) *InMemCache[T] {
	return &InMemCache[T]{gocache.New(defaultExp, defaultExp)}
}

func (c *InMemCache[T]) Set(ctx context.Context, key string, value T) error {
	c.gocache.Set(key, value, gocache.DefaultExpiration)
	return nil
}

func (c *InMemCache[T]) Get(ctx context.Context, key string) (T, error) {
	var value T

	element, found := c.gocache.Get(key)
	if !found {
		return value, fmt.Errorf("not found element with key: %s", key)
	}

	return element.(T), nil
}
