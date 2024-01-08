package service

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/fs"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
)

type UserService interface {
	Update(userDto *dto.UserUpdateRequest) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository, fs *fs.MinioImpl) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) Update(ud *dto.UserUpdateRequest) error {
	_, err := s.repository.Read(ud.ID)
	if err != nil {
		return err
	}

	return nil
}
