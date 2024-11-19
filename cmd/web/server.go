package main

import (
	"github.com/d3cie/pubnode/interface/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func (a *app) startServer() error {
	a.logger.Info("starting server on port " + a.cfg.Port)

	fiberApp := fiber.New(fiber.Config{
		Views:                   html.New("assets/templates", ".html"),
		ViewsLayout:             "layouts/root",
		EnableTrustedProxyCheck: true,
		PassLocalsToViews:       true,
		DisableStartupMessage:   true,
	})

	if a.cfg.Dev {
		fiberApp.Static(
			"/pub",
			"dist",
		)
	}
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("AssetPath", a.cfg.AssetPath)
		return c.Next()
	})

	web.SetupRoutes(fiberApp, a.logger)

	return fiberApp.Listen(":" + a.cfg.Port)
}
