package controllers

import (
	"time"

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
	time.Sleep(4 * time.Second)
	return ctx.SendString("error example")
}
