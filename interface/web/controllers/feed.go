package controllers

import "github.com/gofiber/fiber/v2"

type FeedController struct {
}

func (c *FeedController) Feed_Get(ctx *fiber.Ctx) error {
	return ctx.Render("pages/feed", fiber.Map{
		"ActiveRoute": "feed",
		"User": fiber.Map{
			"Username": "d3cie",
			"Avatar":   "https://avatars.githubusercontent.com/u/36079094?v=4",
		},
	}, "layouts/root", "layouts/main")
}
