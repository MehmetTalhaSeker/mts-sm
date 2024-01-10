package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/service"
)

func AuthRouter(app fiber.Router, service service.AuthService) {
	app.Post("/auth/register", register(service))
	app.Post("/auth/login", login(service))
	app.Delete("/auth/logout", logout(service))
}

func register(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.RegisterRequest)

		err := utils.ParseBodyAndValidate(c, requestBody)
		if err != nil {
			return err
		}

		err = service.Register(requestBody)
		if err != nil {
			return err
		}

		return c.Status(http.StatusCreated).JSON("Proceed to login page.")
	}
}

func login(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.LoginRequest)

		err := utils.ParseBodyAndValidate(c, requestBody)
		if err != nil {
			return err
		}

		authResponse, err := service.Login(requestBody)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(&authResponse)
	}
}

func logout(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ah := c.Get("Authorization")
		if ah == "" {
			return errorutils.ErrUnauthorized
		}

		if err := service.Logout(ah); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON("OK")
	}
}
