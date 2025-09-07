package middlewares

import (
	"strings"

	"blogger/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from cookie (Web) or Authorization header (API)
		token := c.Cookies("token")
		if token == "" {
			authHeader := c.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		c.Locals("user", claims["username"])
		return c.Next()
	}
}
