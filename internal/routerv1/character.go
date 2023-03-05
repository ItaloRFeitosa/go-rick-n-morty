package routerv1

import "github.com/italorfeitosa/go-rick-n-morty/internal/di"

func Characters(c *di.Container) {
	c.FiberApp.Get("/characters", c.CharacterController().SearchCharacters)
}
