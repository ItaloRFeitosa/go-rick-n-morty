package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-rick-n-morty/internal/character"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/process"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	character.Setup(app)

	go app.Listen(fmt.Sprintf(":%s", viper.GetString("PORT")))

	process.GracefulShutdown(func(ctx context.Context) {
		err := app.Server().Shutdown()
		if err != nil {
			log.Println("error on shutdown server: ", err)
		}

	}, viper.GetDuration("GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS"))
}
