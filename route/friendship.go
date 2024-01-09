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

func FriendshipRouter(app fiber.Router, as service.AuthService, service service.FriendshipService) {
	app.Post("/friendships", middleware.AuthMiddleware(as), createFriendship(service))
	app.Put("/friendships", middleware.AuthMiddleware(as), updateFriendship(service))
	app.Delete("/friendships/:id", middleware.AuthMiddleware(as), deleteFriendship(service))
}

func createFriendship(service service.FriendshipService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.FriendshipCreateRequest)

		sid, ok := c.Locals("UserID").(string)
		if !ok {
			return errorutils.ErrInvalidRequest
		}

		i, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}

		err = utils.ParseBody(c, rb)
		if err != nil {
			return err
		}

		rb.UserID = uint(i)

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

		sid, ok := c.Locals("UserID").(string)
		if !ok {
			return errorutils.ErrInvalidRequest
		}

		i, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}

		err = utils.ParseBody(c, rb)
		if err != nil {
			return err
		}

		rb.UserID = uint(i)

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

		sid, ok := c.Locals("UserID").(string)
		if !ok {
			return errorutils.ErrInvalidRequest
		}

		userID, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}
		rb.UserID = uint(userID)

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
