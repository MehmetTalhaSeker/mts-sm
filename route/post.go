package route

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/apputils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/middleware"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func PostRouter(app fiber.Router, as service.AuthService, service service.PostService, cac *cache.Cache) {
	app.Post("/posts", middleware.AuthMiddleware(as), createPost(service))
	app.Get("/posts/:id", readPost(service))
	app.Get("/posts", readsPost(service))
	app.Put("/posts", middleware.AuthMiddleware(as), updatePost(service))
	app.Delete("/posts/:id", middleware.AuthMiddleware(as), deletePost(service))

	app.Post("/posts/shorten/:id", createShortUrl(cac))
}

func createPost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rb dto.PostCreateRequest

		file, err := c.FormFile("photo")
		if err != nil {
			return err
		}

		text := c.FormValue("text")

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

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

func readsPost(service service.PostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, err := service.Reads()
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

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

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

		uid, err := utils.ExtractUserID(c)
		if err != nil {
			return err
		}

		rb.UserID = uid

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

func createShortUrl(cac *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return errorutils.ErrInvalidRequest
		}

		ou := fmt.Sprintf("http://localhost:8080/v1/posts/%s", id)
		suk := apputils.GenerateShortKey()

		cac.Set(suk, ou, time.Minute*60)

		su := fmt.Sprintf("http://localhost:8080/short/%s", suk)

		return c.Status(http.StatusOK).JSON(su)
	}
}
