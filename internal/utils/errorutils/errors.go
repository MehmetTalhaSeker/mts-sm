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
	ErrUnauthorized            = NewError(http.StatusUnauthorized, "Unauthorized user.")
	ErrUsernameAlreadyTaken    = NewError(http.StatusBadRequest, "Username already taken.")
	ErrEmailAlreadyTaken       = NewError(http.StatusBadRequest, "Email already taken.")
	ErrPasswordNotAcceptable   = NewError(http.StatusBadRequest, "Password not acceptable.")
	ErrInvalidPasswordUsername = NewError(http.StatusBadRequest, "Username or password invalid.")
	ErrUserNotFound            = NewError(http.StatusBadRequest, "User not found.")
	ErrInvalidId               = NewError(http.StatusBadRequest, "Invalid ID.")
	ErrInvalidRole             = NewError(http.StatusBadRequest, "Invalid role.")
	ErrEmptyId                 = NewError(http.StatusBadRequest, "ID can't be empty")
	ErrInvalidRequest          = NewError(http.StatusBadRequest, "Invalid request.")
	ErrNotSupportedImage       = NewError(http.StatusBadRequest, "This extension is not supported.")
	ErrFailedSave              = NewError(http.StatusServiceUnavailable, "We couldn't save your request. Please try again!")
	ErrFailedRead              = NewError(http.StatusServiceUnavailable, "We couldn't read your request. Please try again!")
	ErrInvalidQueryParam       = NewError(http.StatusBadRequest, "Your requested query params are invalid. Please try again.")
)
