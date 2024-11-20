package main

import (
	"github.com/d3cie/pubnode/interface/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func (a *app) startServer() error {
	a.logger.Info("starting server on port " + a.cfg.Port)

	//TODO: use embed instead of static
	engine := html.New("templates", ".html")

	fiberApp := fiber.New(fiber.Config{
		Views:                   engine,
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
		fiberApp.Static(
			"/pub",
			"assets/static",
		)
	}
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("AssetPath", a.cfg.AssetPath)
		//TODO: get the version from the proper place?
		c.Locals("AppVersion", "0.0.1")
		return c.Next()
	})

	web.SetupRoutes(fiberApp, a.logger)

	return fiberApp.Listen(":" + a.cfg.Port)
}
