package repository

import (
	"gorm.io/gorm"

	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
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
