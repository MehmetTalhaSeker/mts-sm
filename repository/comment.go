package repository

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(u *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(u *model.Comment) error {
	return r.db.Create(&u).Error
}
