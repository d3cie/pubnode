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

	app.Use(loghttp.New(logger))

	app.Get("/", homeController.Home_Get)
	app.Get("/feed", feedController.Feed_Get)
}
