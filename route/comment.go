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

func CommentRouter(app fiber.Router, as service.AuthService, service service.CommentService) {
	app.Post("/comments", middleware.AuthMiddleware(as), createPostComment(service))
}

func createPostComment(service service.CommentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.CommentCreateRequest)

		sid, ok := c.Locals("UserID").(string)
		if !ok {
			return errorutils.ErrInvalidRequest
		}

		userID, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}
		rb.UserID = uint(userID)

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
