package cache

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

type Cache[T any] interface {
	Get(ctx context.Context, key string) (T, error)
	Set(ctx context.Context, key string, value T) error
}

// Simple Factory Example

func New[T any](exp time.Duration) Cache[T] {
	if viper.GetBool("IS_DEVELOPMENT") {
		return NewInMemCache[T](exp)
	}

	return NewRedisCache[T](exp)
}
