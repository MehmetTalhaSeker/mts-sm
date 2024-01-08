package service

import (
	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/jwt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/password"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AuthService interface {
	Register(register *dto.RegisterDTO) error
	Login(login *dto.LoginDTO) (*dto.AuthResponse, error)
	ReadUser(username string) (*model.User, error)
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (s *authService) Register(rd *dto.RegisterDTO) error {
	um := &model.User{
		Username: rd.Username,
		Password: password.Generate(rd.Password),
	}

	err := s.repository.Create(um)
	if err != nil {
		return errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *authService) Login(login *dto.LoginDTO) (*dto.AuthResponse, error) {
	u, err := s.repository.ReadBy("username", login.Username)
	if err != nil {
		return nil, errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := password.Verify(u.Password, login.Password); err != nil {
		return nil, errorutils.ErrInvalidPasswordUsername
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
	})

	return &dto.AuthResponse{
		Token: t,
	}, nil
}

func (s *authService) ReadUser(username string) (*model.User, error) {
	u, err := s.repository.ReadBy("username", username)
	if err != nil {
		return nil, errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}
	return u, nil
}
