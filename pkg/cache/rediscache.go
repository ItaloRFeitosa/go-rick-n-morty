package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisCache[T any] struct {
	client         *redis.Client
	expirationTime time.Duration
}

func NewRedisCache[T any](expirationTime time.Duration) *RedisCache[T] {
	opt, err := redis.ParseURL(viper.GetString("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)

	statuscmd := client.Ping(context.Background())
	if err := statuscmd.Err(); err != nil {
		log.Fatal(err)
	}

	return &RedisCache[T]{client, expirationTime}
}

func (rc *RedisCache[T]) Get(ctx context.Context, key string) (T, error) {
	var v T

	statuscmd := rc.client.Get(ctx, key)

	if err := statuscmd.Err(); err != nil {
		return v, err
	}

	valueAsString, err := statuscmd.Result()
	if err != nil {
		return v, err
	}

	if err := json.Unmarshal([]byte(valueAsString), &v); err != nil {
		return v, err
	}

	return v, nil
}

func (rc *RedisCache[T]) Set(ctx context.Context, key string, value T) error {
	valueAsString, err := json.Marshal(value)
	if err != nil {
		return err
	}

	statuscmd := rc.client.Set(ctx, key, valueAsString, rc.expirationTime)

	return statuscmd.Err()
}
