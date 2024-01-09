package repository

import (
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"gorm.io/gorm"
)

type FriendshipRepository interface {
	Create(*model.Friendship) error
	ReadBy(field string, value string) (*model.Friendship, error)
	Update(*model.Friendship) (*model.Friendship, error)
	Delete(us *model.Friendship) error
}

type friendshipRepository struct {
	db *gorm.DB
}

func NewFriendshipRepository(db *gorm.DB) FriendshipRepository {
	return &friendshipRepository{
		db: db,
	}
}

func (r *friendshipRepository) Create(u *model.Friendship) error {
	return r.db.Create(&u).Error
}

func (r *friendshipRepository) ReadBy(field string, value string) (*model.Friendship, error) {
	var u model.Friendship

	if err := r.db.Where(fmt.Sprintf("%s = ?", field), value).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *friendshipRepository) Update(us *model.Friendship) (*model.Friendship, error) {
	if err := r.db.Save(&us).Error; err != nil {
		return nil, err
	}

	return us, nil
}

func (r *friendshipRepository) Delete(us *model.Friendship) error {
	tx := r.db.Delete(&us)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
