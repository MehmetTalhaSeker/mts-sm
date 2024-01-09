package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func CommentRouter(app fiber.Router, as service.AuthService, service service.CommentService) {
	app.Post("/comments", middleware.AuthMiddleware(as), createPostComment(service))
}

func createPostComment(service service.CommentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.CommentCreateRequest)

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

		err = utils.ParseBody(c, rb)
		if err != nil {
			return err
		}

		if err = utils.Validate(c, rb); err != nil {
			return err
		}

		if err = service.Create(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
