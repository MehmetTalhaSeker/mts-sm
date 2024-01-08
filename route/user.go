package route

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func UserRouter(app fiber.Router, as service.AuthService, service service.UserService) {
	app.Put("/users", middleware.AuthMiddleware(as), updateUser(service))
}

func updateUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rb dto.UserUpdateRequest

		file, err := c.FormFile("photo")
		if err != nil {
			return err
		}

		sid, ok := c.Locals("UserID").(string)

		if !ok {
			return errorutils.ErrInvalidRequest
		}

		i, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}

		rb.ID = uint(i)
		rb.Photo = file

		if err = utils.Validate(c, &rb); err != nil {
			return err
		}

		if err = service.Update(&rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
