package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-rick-n-morty/internal/character"
)

func main() {
	app := fiber.New()

	character.Setup(app)

	app.Listen(":8080")
}
