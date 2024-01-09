package service

import (
	"fmt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/types"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type FriendshipService interface {
	Create(*dto.FriendshipCreateRequest) error
	Update(*dto.FriendshipUpdateRequest) error
	Delete(*dto.FriendshipDeleteRequest) error
}

type friendshipService struct {
	repository repository.FriendshipRepository
	conf       *config.Config
}

func NewFriendshipService(repository repository.FriendshipRepository, conf *config.Config) FriendshipService {
	return &friendshipService{
		repository: repository,
		conf:       conf,
	}
}

func (s *friendshipService) Create(ud *dto.FriendshipCreateRequest) error {
	var id string

	if ud.UserID > ud.TargetUserID {
		id = fmt.Sprintf("%s-%s", strconv.Itoa(int(ud.UserID)), strconv.Itoa(int(ud.TargetUserID)))
	} else {
		id = fmt.Sprintf("%s-%s", strconv.Itoa(int(ud.TargetUserID)), strconv.Itoa(int(ud.UserID)))
	}

	um := &model.Friendship{
		TargetUserID: ud.TargetUserID,
		UserID:       ud.UserID,
		Status:       types.Pending,
		ID:           id,
	}

	err := s.repository.Create(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *friendshipService) Update(ud *dto.FriendshipUpdateRequest) error {
	u, err := s.repository.ReadBy("id", ud.ID)
	if err != nil {
		return err
	}

	if u.TargetUserID != ud.UserID {
		return errorutils.ErrUnauthorized
	}

	u.Status = ud.Status

	_, err = s.repository.Update(u)
	if err != nil {
		return errorutils.ErrFailedSave
	}

	return nil
}

func (s *friendshipService) Delete(ud *dto.FriendshipDeleteRequest) error {
	u, err := s.repository.ReadBy("id", ud.ID)
	if err != nil {
		return err
	}

	if u.UserID != ud.UserID && u.TargetUserID != ud.UserID {
		return errorutils.ErrUnauthorized
	}

	err = s.repository.Delete(u)
	if err != nil {
		return errorutils.ErrInvalidRequest
	}

	return nil
}
