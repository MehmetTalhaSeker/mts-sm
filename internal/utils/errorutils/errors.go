package errorutils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e HttpError) Error() string {
	return e.Message
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	status := http.StatusInternalServerError

	if e, ok := err.(*HttpError); ok {
		status = e.Status
	}

	return c.Status(status).JSON(&HttpError{
		Status:  status,
		Message: err.Error(),
	})
}

func NewError(status int, message string) *HttpError {
	return &HttpError{
		Status:  status,
		Message: message,
	}
}

var (
	ErrUpdateTimeExpired       = NewError(http.StatusBadRequest, "Update time expired.")
	ErrUnauthorized            = NewError(http.StatusUnauthorized, "Unauthorized user.")
	ErrInvalidPasswordUsername = NewError(http.StatusBadRequest, "Username or password invalid.")
	ErrInvalidRequest          = NewError(http.StatusBadRequest, "Invalid request.")
	ErrNotSupportedImage       = NewError(http.StatusBadRequest, "This extension is not supported.")
	ErrFailedSave              = NewError(http.StatusServiceUnavailable, "We couldn't save your request. Please try again!")
)
