package service

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/fs"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
	"path/filepath"
	"time"
)

type PostService interface {
	Create(*dto.PostCreateRequest) error
	Read(ID uint) (*model.Post, error)
	Update(*dto.PostUpdateRequest) error
	Delete(*dto.DeleteRequest) error
}

type postService struct {
	repository repository.PostRepository
	fs         fs.IFileStorage
	conf       *config.Config
}

func NewPostService(repository repository.PostRepository, fs fs.IFileStorage, conf *config.Config) PostService {
	return &postService{
		repository: repository,
		fs:         fs,
		conf:       conf,
	}
}

func (s *postService) Create(ud *dto.PostCreateRequest) error {
	var imageUrl *string

	if ud.Photo != nil {
		fileExtension := filepath.Ext(ud.Photo.Filename)
		if !utils.Contains(s.conf.Minio.SupportedExtensions, fileExtension) {
			return errorutils.ErrNotSupportedImage
		}

		open, err := ud.Photo.Open()
		imageUrl, err = s.fs.UploadImage("postPictures", fileExtension, open, ud.Photo.Size)
		if err != nil {
			return errorutils.ErrFailedSave
		}
	}

	um := &model.Post{
		Text:   ud.Text,
		Image:  *imageUrl,
		UserID: ud.UserID,
	}

	err := s.repository.Create(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *postService) Read(ID uint) (*model.Post, error) {
	u, err := s.repository.Read(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *postService) Update(ud *dto.PostUpdateRequest) error {
	u, err := s.repository.Read(ud.ID)
	if err != nil {
		return err
	}

	if u.UserID != ud.UserID {
		return errorutils.ErrUnauthorized
	}

	if time.Since(u.CreatedAt) > 5*time.Minute {
		return errorutils.ErrUpdateTimeExpired
	}

	var imageUrl *string

	if ud.Photo != nil {
		fileExtension := filepath.Ext(ud.Photo.Filename)
		if !utils.Contains(s.conf.Minio.SupportedExtensions, fileExtension) {
			return errorutils.ErrNotSupportedImage
		}

		open, err := ud.Photo.Open()
		imageUrl, err = s.fs.UploadImage("postPictures", fileExtension, open, ud.Photo.Size)
		if err != nil {
			return errorutils.ErrFailedSave
		}
	}

	if imageUrl != nil {
		u.Image = *imageUrl
	}

	if ud.Text != "" {
		u.Text = ud.Text
	}

	if imageUrl == nil && ud.Text == "" {
		return errorutils.ErrInvalidRequest
	}

	u.UpdatedAt = time.Now()

	_, err = s.repository.Update(u)
	if err != nil {
		return errorutils.ErrFailedSave
	}

	return nil
}

func (s *postService) Delete(ud *dto.DeleteRequest) error {
	u, err := s.repository.Read(ud.ID)
	if err != nil {
		return err
	}

	if u.UserID != ud.UserID {
		return errorutils.ErrUnauthorized
	}

	err = s.repository.Delete(u.ID)
	if err != nil {
		return errorutils.ErrInvalidRequest
	}

	return nil
}
