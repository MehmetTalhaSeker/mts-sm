package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
)

type LikeService interface {
	CreatePostLike(*dto.PostLikeCreateRequest) error
	CreateCommentLike(*dto.CommentLikeCreateRequest) error
}

type likeService struct {
	repository repository.LikeRepository
	conf       *config.Config
}

func NewLikeService(repository repository.LikeRepository, conf *config.Config) LikeService {
	return &likeService{
		repository: repository,
		conf:       conf,
	}
}

func (s *likeService) CreatePostLike(ud *dto.PostLikeCreateRequest) error {
	um := &model.PostLike{
		PostID:    ud.PostID,
		UserID:    ud.UserID,
		ID:        fmt.Sprintf("%s-%s", strconv.Itoa(int(ud.PostID)), strconv.Itoa(int(ud.UserID))),
		CreatedAt: time.Now(),
	}

	err := s.repository.CreatePostLike(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *likeService) CreateCommentLike(ud *dto.CommentLikeCreateRequest) error {
	um := &model.CommentLike{
		CommentID: ud.CommentID,
		UserID:    ud.UserID,
		ID:        fmt.Sprintf("%s-%s", strconv.Itoa(int(ud.CommentID)), strconv.Itoa(int(ud.UserID))),
		CreatedAt: time.Now(),
	}

	err := s.repository.CreateCommentLike(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
