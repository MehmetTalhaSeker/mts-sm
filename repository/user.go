package repository

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
)

type UserRepository interface {
	Create(*model.User) error
	ReadBy(field string, value string) (*model.User, error)
	Read(id uint) (*model.User, error)
	Update(*model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(u *model.User) error {
	res := r.db.Create(&u)

	var pgErr *pgconn.PgError
	if errors.As(res.Error, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return errors.New("username in use")
		}
	}

	return res.Error
}

func (r *userRepository) Read(id uint) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) ReadBy(field string, value string) (*model.User, error) {
	var u model.User

	if err := r.db.Where(fmt.Sprintf("%s = ?", field), value).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) Update(us *model.User) (*model.User, error) {
	if err := r.db.Save(&us).Error; err != nil {
		return nil, err
	}

	return us, nil
}
