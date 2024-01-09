package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Image  string
	UserID uint
	Text   string `gorm:"type:text;character set:utf8mb4;collate:utf8mb4_unicode_ci"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
