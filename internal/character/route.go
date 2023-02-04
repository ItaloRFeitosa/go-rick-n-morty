package character

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	c := setupController()
	app.Get("/characters", c.SearchCharacters)
}
