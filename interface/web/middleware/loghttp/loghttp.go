package loghttp

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New(logger *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		userAttrs := slog.Group("user", slog.String("ip", c.IP()), slog.String("user_agent", c.Get("User-Agent")))
		requestAttrs := slog.Group("request", slog.String("method", c.Method()), slog.String("path", c.Path()))
		responseAttrs := slog.Group("repsonse", slog.Int("status", c.Response().StatusCode()), slog.String("duration", fmt.Sprintf("%dms", -time.Until(start).Milliseconds())))

		logger.Info("http", userAttrs, requestAttrs, responseAttrs)
		return err
	}
}
