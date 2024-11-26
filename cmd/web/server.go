package main

import (
	"net/http"

	"github.com/d3cie/pubnode/interface/web/controllers"
	"github.com/d3cie/pubnode/interface/web/middleware/loghttp"
	"github.com/d3cie/pubnode/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func (a *app) startServer() error {
	a.logger.Info("starting server on port " + a.cfg.Port)
	engine := html.NewFileSystem(http.FS(templates.EmbeddedFiles), ".html")

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
			"static",
		)
	}
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("Config", map[string]string{
			"AppVersion": a.cfg.AppVersion,
			"AssetPath":  a.cfg.AssetPath,
		})
		c.Locals("AppVersion", a.cfg.AppVersion)
		return c.Next()
	})

	homeController := controllers.HomeController{}
	feedController := controllers.FeedController{}
	postsController := controllers.PostsController{}
	authController := controllers.AuthController{}

	fiberApp.Use(loghttp.New(a.logger))

	fiberApp.Get("/", homeController.Home_Get)
	fiberApp.Get("/login", authController.Login_Get)
	fiberApp.Post("/login", authController.Login_Post)
	fiberApp.Get("/feed", feedController.Feed_Get)
	fiberApp.Get("/post/new", postsController.NewPost_Get)

	return fiberApp.Listen(":" + a.cfg.Port)
}
