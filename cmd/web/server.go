package main

import (
	"github.com/d3cie/pubnode/interface/web/routes"
	"github.com/gofiber/fiber/v2"
)

func (a *app) startServer() error {
	a.logger.Info("starting server on port: " + a.cfg.Port)

	fiberApp := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		DisableStartupMessage:   true,
	})

	// initialize dependencies
	routes.SetupHomeRoutes(fiberApp)

	return fiberApp.Listen(":" + a.cfg.Port)
}
