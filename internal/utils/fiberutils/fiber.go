package utils

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/validatorutils"
	"github.com/gofiber/fiber/v2"
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
