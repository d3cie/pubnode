package web

import (
	"log/slog"

	"github.com/d3cie/pubnode/interface/web/routes"
	"github.com/d3cie/pubnode/internal/infra/db"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *db.DB, logger *slog.Logger) {
	routes.SetupHomeRoutes(app)
}
