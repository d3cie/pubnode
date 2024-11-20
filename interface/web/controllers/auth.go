package controllers

import "github.com/gofiber/fiber/v2"

type AuthController struct {
}

func (c *AuthController) Login_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/login", fiber.Map{}, "layouts/root")
}
