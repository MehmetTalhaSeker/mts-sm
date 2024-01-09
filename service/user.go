package service

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/fs"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"gorm.io/gorm/utils"
	"path/filepath"
	"time"
)

type UserService interface {
	Update(userDto *dto.UserUpdateRequest) error
}

type userService struct {
	repository repository.UserRepository
	fs         fs.IFileStorage
	conf       *config.Config
}

func NewUserService(repository repository.UserRepository, fs fs.IFileStorage, conf *config.Config) UserService {
	return &userService{
		repository: repository,
		fs:         fs,
		conf:       conf,
	}
}

func (s *userService) Update(ud *dto.UserUpdateRequest) error {
	u, err := s.repository.Read(ud.ID)
	if err != nil {
		return err
	}

	fileExtension := filepath.Ext(ud.Photo.Filename)
	if !utils.Contains(s.conf.Minio.SupportedExtensions, fileExtension) {
		return errorutils.ErrNotSupportedImage
	}

	open, err := ud.Photo.Open()
	imageUrl, err := s.fs.UploadImage("profilePictures", fileExtension, open, ud.Photo.Size)
	if err != nil {
		return errorutils.ErrFailedSave
	}

	u.ProfilePicture = *imageUrl
	u.UpdatedAt = time.Now()

	_, err = s.repository.Update(u)
	if err != nil {
		return errorutils.ErrFailedSave
	}

	return nil
}

func (s *userService) AddFriend() error {

	return nil
}
