package middleware

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/jwt"
	"github.com/MehmetTalhaSeker/mts-sm/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthMiddleware(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := c.Get("Authorization")
		if h == "" {
			return errorutils.ErrUnauthorized
		}

		chunks := strings.Split(h, " ")

		if len(chunks) < 2 {
			return errorutils.ErrUnauthorized
		}

		tokenPayload, err := jwt.Verify(chunks[1])

		if err != nil {
			return errorutils.ErrUnauthorized
		}

		_, err = service.ReadUser(tokenPayload.Username)
		if err != nil {
			return errorutils.ErrUnauthorized
		}

		c.Locals("UserID", tokenPayload.ID)
		c.Locals("Username", tokenPayload.Username)

		return c.Next()
	}
}
