package validatorutils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
)

var validate = validator.New()

func Validate(payload interface{}) error {
	err := validate.Struct(payload)
	if err != nil {
		var errors []string

		for _, err := range err.(validator.ValidationErrors) {
			message := ""

			l := 0
			switch err.Value().(type) {
			case string:
				l = len(err.Value().(string))
			case []string:
				l = len(err.Value().([]string))
			}

			if l == 0 {
				message = fmt.Sprintf("%v is required.", err.Field())
			} else if err.Tag() == "max" {
				message = fmt.Sprintf("%v character count should be less then %v.", err.Field(), err.Param())
			} else if err.Tag() == "min" {
				message = fmt.Sprintf("%v character count should be more then %v.", err.Field(), err.Param())
			} else {
				message = fmt.Sprintf("%v doesn't satisfy the constraint", err.Field())
			}

			errors = append(
				errors,
				message,
			)
		}

		return &errorutils.HttpError{
			Status:  http.StatusBadRequest,
			Message: strings.Join(errors, "|"),
		}
	}

	return nil
}

var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l < 34
})
