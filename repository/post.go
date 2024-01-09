package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
)

type PostRepository interface {
	Create(*model.Post) error
	ReadBy(field string, value string) (*model.Post, error)
	Read(id uint) (*model.Post, error)
	Update(*model.Post) (*model.Post, error)
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(u *model.Post) error {
	return r.db.Create(&u).Error
}

func (r *postRepository) Read(id uint) (*model.Post, error) {
	var u model.Post
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *postRepository) ReadBy(field string, value string) (*model.Post, error) {
	var u model.Post

	if err := r.db.Where(fmt.Sprintf("%s = ?", field), value).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *postRepository) Update(us *model.Post) (*model.Post, error) {
	if err := r.db.Save(&us).Error; err != nil {
		return nil, err
	}

	return us, nil
}

func (r *postRepository) Delete(id uint) error {
	tx := r.db.Delete(&model.Post{}, id)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
