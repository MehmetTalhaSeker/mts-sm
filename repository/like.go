package repository

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreatePostLike(u *model.PostLike) error
	CreateCommentLike(u *model.CommentLike) error
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{
		db: db,
	}
}

func (r *likeRepository) CreatePostLike(u *model.PostLike) error {
	return r.db.Create(&u).Error
}

func (r *likeRepository) CreateCommentLike(u *model.CommentLike) error {
	return r.db.Create(&u).Error
}
