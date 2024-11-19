package routes

import "github.com/gofiber/fiber/v2"

type homeRoutesController struct {
}

func SetupHomeRoutes(app *fiber.App) {
	c := homeRoutesController{}
	app.Get("/", c.index)
}

func (c *homeRoutesController) index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
