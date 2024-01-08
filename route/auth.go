package route

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	utils "github.com/MehmetTalhaSeker/mts-sm/internal/utils/fiberutils"
	"github.com/MehmetTalhaSeker/mts-sm/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthRouter(app fiber.Router, service service.AuthService) {
	app.Post("/auth/register", register(service))
	app.Post("/auth/login", login(service))
}

func register(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.RegisterDTO)
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
		requestBody := new(dto.LoginDTO)
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
