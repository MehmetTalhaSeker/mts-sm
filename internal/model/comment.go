package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ParentID *uint
	Parent   *Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PostID   uint     `gorm:"not null"`
	Post     Post     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID   uint     `gorm:"not null"`
	User     User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Text     string   `gorm:"type:text;character set:utf8mb4;collate:utf8mb4_unicode_ci"`
}
