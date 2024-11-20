package web

import (
	"log/slog"

	"github.com/d3cie/pubnode/interface/web/controllers"
	"github.com/d3cie/pubnode/interface/web/middleware/loghttp"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, logger *slog.Logger) {
	homeController := controllers.HomeController{}
	feedController := controllers.FeedController{}
	postsController := controllers.PostsController{}
	authController := controllers.AuthController{}

	app.Use(loghttp.New(logger))

	app.Get("/", homeController.Home_Get)
	app.Get("/login", authController.Login_Get)
	app.Get("/feed", feedController.Feed_Get)
	app.Get("/create", postsController.Create_Get)
}
