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

func PostRouter(app fiber.Router, as service.AuthService, service service.PostService) {
	app.Post("/posts", middleware.AuthMiddleware(as), createPost(service))
	app.Get("/posts/:id", readPost(service))
	app.Put("/posts", middleware.AuthMiddleware(as), updatePost(service))
	app.Delete("/posts/:id", middleware.AuthMiddleware(as), deletePost(service))
}

func createPost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rb dto.PostCreateRequest

		file, err := c.FormFile("photo")
		if err != nil {
			return err
		}

		text := c.FormValue("text")
		sid, ok := c.Locals("UserID").(string)

		if !ok {
			return errorutils.ErrInvalidRequest
		}

		i, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}

		rb.UserID = uint(i)
		rb.Photo = file
		rb.Text = text

		if err = utils.Validate(c, &rb); err != nil {
			return err
		}

		if err = service.Create(&rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func readPost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ID, err := strconv.Atoi(id)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}

		u, err := service.Read(uint(ID))
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(u)
	}
}

func updatePost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rb dto.PostUpdateRequest

		file, err := c.FormFile("photo")
		if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
			return err
		}

		text := c.FormValue("text")
		id := c.FormValue("id")
		sid, ok := c.Locals("UserID").(string)

		if !ok {
			return errorutils.ErrInvalidRequest
		}

		userID, err := strconv.Atoi(sid)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}
		rb.UserID = uint(userID)

		ID, err := strconv.Atoi(id)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}
		rb.ID = uint(ID)

		rb.Photo = file
		rb.Text = text

		if err = utils.Validate(c, &rb); err != nil {
			return err
		}

		if err = service.Update(&rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}

func deletePost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rb := new(dto.DeleteRequest)

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
		ID, err := strconv.Atoi(id)
		if err != nil {
			return errorutils.ErrInvalidRequest
		}
		rb.ID = uint(ID)

		if err = utils.Validate(c, rb); err != nil {
			return err
		}

		if err = service.Delete(rb); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
