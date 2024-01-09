package dto

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/types"
)

type FriendshipCreateRequest struct {
	TargetUserID uint `json:"targetUserId'" validate:"required"`
	UserID       uint `validate:"required"`
}

type FriendshipUpdateRequest struct {
	ID     string                 `json:"id"           validate:"required"`
	Status types.FriendshipStatus `json:"status"       validate:"required"`
	UserID uint                   `validate:"required"`
}

type FriendshipDeleteRequest struct {
	ID     string `json:"id'"          validate:"required"`
	UserID uint   `validate:"required"`
}
