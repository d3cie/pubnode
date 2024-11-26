package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func (c *AuthController) Login_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/login", fiber.Map{
		"Meta": fiber.Map{
			"Title":       "Login to Pubnode",
			"Description": "Login to continue",
		},
	}, "layouts/root")
}

func (c *AuthController) Login_Post(ctx *fiber.Ctx) error {
	return ctx.SendString("error example")
}

func (c *AuthController) Register_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/register", fiber.Map{
		"Meta": fiber.Map{
			"Title":       "Register on Pubnode",
			"Description": "Register to continue",
		},
	}, "layouts/root")
}

func (c *AuthController) Register_Post(ctx *fiber.Ctx) error {
	return ctx.SendString("error example")
}
