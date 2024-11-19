package controllers

import "github.com/gofiber/fiber/v2"

type HomeController struct {
}

func (c *HomeController) Home_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/landing", fiber.Map{})
}
