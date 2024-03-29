package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"

	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/validatorutils"
)

func ParseBody(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return validatorutils.Validate(body)
}

func Validate(ctx *fiber.Ctx, body interface{}) error {
	return validatorutils.Validate(body)
}

func ExtractUserID(ctx *fiber.Ctx) (uint, error) {
	sid, ok := ctx.Locals("UserID").(string)
	if !ok {
		return 0, errorutils.ErrInvalidRequest
	}

	userID, err := strconv.Atoi(sid)
	if err != nil {
		return 0, errorutils.ErrInvalidRequest
	}

	return uint(userID), nil
}

func ShortenerRedirect(cac *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")
		if key == "" {
			return errorutils.ErrInvalidRequest
		}

		v, b := cac.Get(key)
		if !b {
			return errorutils.ErrInvalidRequest
		}

		url, ok := v.(string)
		if !ok {
			return errorutils.ErrInvalidRequest
		}

		return c.Redirect(url)
	}
}
