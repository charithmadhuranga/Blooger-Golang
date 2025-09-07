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

// JWTAuthRedirect enforces auth for web routes; redirects to /login when missing/invalid
func JWTAuthRedirect() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Cookies("token")
        if token == "" {
            authHeader := c.Get("Authorization")
            if strings.HasPrefix(authHeader, "Bearer ") {
                token = strings.TrimPrefix(authHeader, "Bearer ")
            }
        }

        if token == "" {
            return c.Redirect("/login")
        }

        claims, err := utils.ParseJWT(token)
        if err != nil {
            return c.Redirect("/login")
        }
        c.Locals("user", claims["username"])
        return c.Next()
    }
}

// OptionalUser sets user in locals if a valid token is present; never blocks
func OptionalUser() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Cookies("token")
        if token == "" {
            authHeader := c.Get("Authorization")
            if strings.HasPrefix(authHeader, "Bearer ") {
                token = strings.TrimPrefix(authHeader, "Bearer ")
            }
        }
        if token != "" {
            if claims, err := utils.ParseJWT(token); err == nil {
                c.Locals("user", claims["username"])
            }
        }
        return c.Next()
    }
}
