package character

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-rick-n-morty/pkg/ricknmorty"
)

type Controller interface {
	SearchCharacters(c *fiber.Ctx) error
}

type controller struct {
	ricknmorty ricknmorty.Client
}

func NewController(r ricknmorty.Client) *controller {
	return &controller{r}
}

func (ctl *controller) SearchCharacters(c *fiber.Ctx) error {
	var (
		chars  ricknmorty.PaginatedCharacters
		filter ricknmorty.FilterCharactersQuery
		err    error
	)

	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	chars, err = ctl.ricknmorty.FilterCharacters(c.UserContext(), filter)

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(chars)
}
