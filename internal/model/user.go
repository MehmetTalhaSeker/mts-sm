package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	ProfilePicture string
	Friends        []User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}
