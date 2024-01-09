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

func LikeRouter(app fiber.Router, as service.AuthService, service service.LikeService) {
	app.Post("/post-likes", middleware.AuthMiddleware(as), createPostLike(service))
	app.Post("/comment-likes", middleware.AuthMiddleware(as), createCommentLike(service))
}

func createPostLike(service service.LikeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.PostLikeCreateRequest)

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

		if err = service.CreatePostLike(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func createCommentLike(service service.LikeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.CommentLikeCreateRequest)

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

		if err = service.CreateCommentLike(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
