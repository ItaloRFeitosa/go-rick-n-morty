package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Singleton Pattern Example

var client *redis.Client

func Client() *redis.Client {
	if client != nil {
		return client
	}

	opt, err := redis.ParseURL(viper.GetString("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	client = redis.NewClient(opt)

	statuscmd := client.Ping(context.Background())
	if err := statuscmd.Err(); err != nil {
		log.Fatal(err)
	}

	return client
}
