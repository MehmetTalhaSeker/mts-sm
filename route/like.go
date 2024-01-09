package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func LikeRouter(app fiber.Router, as service.AuthService, service service.LikeService) {
	app.Post("/post-likes", middleware.AuthMiddleware(as), createPostLike(service))
	app.Post("/comment-likes", middleware.AuthMiddleware(as), createCommentLike(service))
}

func createPostLike(service service.LikeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.PostLikeCreateRequest)

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

		if err = service.CreatePostLike(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func createCommentLike(service service.LikeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.CommentLikeCreateRequest)

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

		if err = service.CreateCommentLike(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
