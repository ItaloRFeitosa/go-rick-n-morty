package main

import (
	"context"
	"fmt"
	"log"

	"github.com/italorfeitosa/go-rick-n-morty/internal/di"
	"github.com/italorfeitosa/go-rick-n-morty/internal/routerv1"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/process"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := di.NewContainer()

	routerv1.Characters(c)

	go c.FiberApp.Listen(fmt.Sprintf(":%s", viper.GetString("PORT")))

	process.GracefulShutdown(func(ctx context.Context) {
		err := c.FiberApp.Server().Shutdown()
		if err != nil {
			log.Println("error on shutdown server: ", err)
		}

	}, viper.GetDuration("GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS"))
}
