package service

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
)

type CommentService interface {
	Create(*dto.CommentCreateRequest) error
}

type commentService struct {
	repository repository.CommentRepository
	conf       *config.Config
}

func NewCommentService(repository repository.CommentRepository, conf *config.Config) CommentService {
	return &commentService{
		repository: repository,
		conf:       conf,
	}
}

func (s *commentService) Create(ud *dto.CommentCreateRequest) error {
	um := &model.Comment{
		PostID:   ud.PostID,
		ParentID: ud.ParentID,
		UserID:   ud.UserID,
		Text:     ud.Text,
	}

	err := s.repository.Create(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
