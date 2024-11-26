package controllers

import "github.com/gofiber/fiber/v2"

type PostsController struct {
}

func (c *PostsController) NewPost_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/post/new/text", fiber.Map{
		"Type": "text",
		"User": fiber.Map{
			"Username": "d3cie",
			"Avatar":   "https://avatars.githubusercontent.com/u/36079094?v=4",
		},
	}, "layouts/root", "layouts/post", "layouts/main")
}
