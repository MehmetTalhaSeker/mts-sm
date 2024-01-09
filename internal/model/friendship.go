package model

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/types"
	"time"
)

type Friendship struct {
	UserID       uint   `gorm:"not null"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetUserID uint   `gorm:"not null"`
	TargetUser   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ID           string `gorm:"primarykey"`
	CreatedAt    time.Time
	Status       types.FriendshipStatus
}
