package model

import (
	"gorm.io/gorm"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
)

type Post struct {
	gorm.Model
	Image  string
	UserID uint   `gorm:"not null"`
	Text   string `gorm:"type:text;character set:utf8mb4;collate:utf8mb4_unicode_ci"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Post) ToPublicDTO() *dto.PostResponse {
	return &dto.PostResponse{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Image:     p.Image,
		UserID:    p.UserID,
		Text:      p.Text,
	}
}
