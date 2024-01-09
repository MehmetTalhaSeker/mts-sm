package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
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

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.ID = uid

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
