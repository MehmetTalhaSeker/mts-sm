package dto

import "mime/multipart"

type UserUpdateRequest struct {
	ID    uint                  `validate:"required"`
	Photo *multipart.FileHeader `form:"photo"        validate:"required"`
}
