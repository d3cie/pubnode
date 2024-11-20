package controllers

import "github.com/gofiber/fiber/v2"

type PostsController struct {
}

func (c *PostsController) Create_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/post/link", fiber.Map{
		"User": fiber.Map{
			"Username": "d3cie",
			"Avatar":   "https://avatars.githubusercontent.com/u/36079094?v=4",
		},
	}, "layouts/root", "layouts/create", "layouts/main")
}
