package model

import (
	"time"
)

type PostLike struct {
	Post      Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID    uint   `gorm:"not null"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint   `gorm:"not null"`
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
}

type CommentLike struct {
	Comment   Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentID uint    `gorm:"not null"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint    `gorm:"not null"`
	ID        string  `gorm:"primarykey"`
	CreatedAt time.Time
}
