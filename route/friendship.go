package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func FriendshipRouter(app fiber.Router, as service.AuthService, service service.FriendshipService) {
	app.Post("/friendships", middleware.AuthMiddleware(as), createFriendship(service))
	app.Put("/friendships", middleware.AuthMiddleware(as), updateFriendship(service))
	app.Delete("/friendships/:id", middleware.AuthMiddleware(as), deleteFriendship(service))
}

func createFriendship(service service.FriendshipService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.FriendshipCreateRequest)

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

		err = utils.ParseBody(c, rb)
		if err != nil {
			return err
		}

		rb.UserID = uid

		if err = utils.Validate(c, rb); err != nil {
			return err
		}

		if err = service.Create(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func updateFriendship(service service.FriendshipService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.FriendshipUpdateRequest)

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		err = utils.ParseBody(c, rb)
		if err != nil {
			return err
		}

		rb.UserID = uid

		if err = utils.Validate(c, rb); err != nil {
			return err
		}

		if err = service.Update(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func deleteFriendship(service service.FriendshipService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.FriendshipDeleteRequest)

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

		id := c.Params("id")
		rb.ID = id

		if err = utils.Validate(c, rb); err != nil {
			return err
		}

		if err = service.Delete(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
