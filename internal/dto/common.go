package dto

type DeleteRequest struct {
	ID     uint `json:"id'"          validate:"required"`
	UserID uint `validate:"required"`
}
