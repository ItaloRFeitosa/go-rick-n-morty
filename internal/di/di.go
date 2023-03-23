package di

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-rick-n-morty/internal/character"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/alert"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/cache"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
	"github.com/spf13/viper"
)

type Container struct {
	FiberApp            *fiber.App
	RickNMorty          ricknmorty.Client
	characterController character.Controller

	SimpleHttpAlertManager character.RickNMortyAlertManager
	DiscordAlertManager    character.RickNMortyAlertManager
}

func NewContainer() *Container {
	c := new(Container)

	provideFiberApp(c)
	provideSimpleHttpAlertManager(c)
	provideDiscordAlertManager(c)
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

func provideSimpleHttpAlertManager(c *Container) {
	c.SimpleHttpAlertManager = character.NewRickNMortyClientAlertManager(
		alert.NewSimpleAlertFactory(),
		alert.NewHttpStrategy(viper.GetString("ALERT_URL")),
	)
}

func provideDiscordAlertManager(c *Container) {
	c.DiscordAlertManager = character.NewRickNMortyClientAlertManager(
		alert.NewDiscordAlertFactory("Alert Bot"),
		alert.NewHttpStrategy(viper.GetString("DISCORD_WEBHOOK_URL")),
	)
}

func provideRickNMorty(c *Container) {
	client := ricknmorty.New()

	cacheProxy := character.NewRickNMortyClientCacheProxy(client, cache.New[ricknmorty.PaginatedCharacters](10*time.Second))

	errorDecorator := character.NewRickNMortyClientErrorDecorator(cacheProxy)

	alertDecorator := character.NewRickNMortyClientAlertDecorator(
		errorDecorator,
		c.DiscordAlertManager,
		c.SimpleHttpAlertManager,
	)

	c.RickNMorty = character.NewRickNMortyClientLogDecorator(alertDecorator)
}

func provideCharacterController(c *Container) {
	c.characterController = character.NewController(c.RickNMorty)
}
