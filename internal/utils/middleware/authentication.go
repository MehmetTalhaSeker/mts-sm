package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/jwt"
	"github.com/MehmetTalhaSeker/mts-sm/service"
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

		if err := service.CheckTokenValidity(chunks[1]); err != nil {
			return errorutils.ErrUnauthorized
		}

		tokenPayload, err := jwt.Verify(chunks[1], service.GetJWTKey())
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
