package di

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-rick-n-morty/internal/character"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/cache"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
)

type Container struct {
	FiberApp            *fiber.App
	RickNMorty          ricknmorty.Client
	characterController character.Controller
}

func NewContainer() *Container {
	c := new(Container)

	provideFiberApp(c)
	provideRickNMorty(c)

	return c
}

func (c *Container) CharacterController() character.Controller {
	if c.characterController == nil {
		provideCharacterController(c)
	}

	return c.characterController
}

func provideFiberApp(c *Container) {
	c.FiberApp = fiber.New()
}

func provideRickNMorty(c *Container) {
	client := ricknmorty.New()

	cacheProxy := character.NewRickNMortyClientCacheProxy(client, cache.New[ricknmorty.PaginatedCharacters](10*time.Second))

	alertDecorator := character.NewRickNMortyClientAlertDecorator(cacheProxy)

	c.RickNMorty = character.NewRickNMortyClientLogDecorator(alertDecorator)
}

func provideCharacterController(c *Container) {
	c.characterController = character.NewController(c.RickNMorty)
}
