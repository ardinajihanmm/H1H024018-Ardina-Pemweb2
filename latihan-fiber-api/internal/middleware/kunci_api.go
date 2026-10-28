package middleware

import (
	"latihan-fiber-api/internal/config"

	"github.com/gofiber/fiber/v3"
)

func KunciApi() fiber.Handler {
	konf := config.Muat()
	apiKey := konf.APIKey

	return func(c fiber.Ctx) error {
		kunci := c.Get("X-Api-Key")

		if apiKey == "" || kunci != apiKey {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"API key tidak valid atau tidak tersedia",
			)
		}

		return c.Next()
	}
}