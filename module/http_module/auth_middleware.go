package http_module

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		_ = strings.TrimPrefix(authHeader, "Bearer ")
		//if !validateToken(token) {
		//	return c.Status(401).SendString("Invalid token")
		//}
	} else if strings.HasPrefix(authHeader, "DSN ") {
		_ = strings.TrimPrefix(authHeader, "DSN ")
		//if !validateDSN(dsn) {
		//	return c.Status(401).SendString("Invalid DSN")
		//}
	} else {
		return c.Status(401).SendString("Missing auth")
	}
	return c.Next()
}
