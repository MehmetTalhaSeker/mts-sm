package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"

	"github.com/MehmetTalhaSeker/mts-sm/internal/dto"
	"github.com/MehmetTalhaSeker/mts-sm/internal/model"
	"github.com/MehmetTalhaSeker/mts-sm/internal/shared/config"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/errorutils"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/jwt"
	"github.com/MehmetTalhaSeker/mts-sm/internal/utils/password"
	"github.com/MehmetTalhaSeker/mts-sm/repository"
)

type AuthService interface {
	Register(register *dto.RegisterRequest) error
	Login(login *dto.LoginRequest) (*dto.AuthResponse, error)
	ReadUser(username string) (*model.User, error)
	Logout(header string) error
	CheckTokenValidity(token string) error
	GetJWTKey() string
	GetJWTExp() string
}

type authService struct {
	repository repository.UserRepository
	conf       *config.Config
	cac        *cache.Cache
}

func NewAuthService(repository repository.UserRepository, conf *config.Config, cac *cache.Cache) AuthService {
	return &authService{
		repository: repository,
		conf:       conf,
		cac:        cac,
	}
}

func (s *authService) Register(rd *dto.RegisterRequest) error {
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

func (s *authService) Login(login *dto.LoginRequest) (*dto.AuthResponse, error) {
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
	}, s.conf.Security.Jwt.Exp, s.conf.Security.Jwt.Key)

	return &dto.AuthResponse{
		Token: t,
	}, nil
}

func (s *authService) Logout(h string) error {
	chunks := strings.Split(h, " ")

	if len(chunks) < 2 {
		return errorutils.ErrUnauthorized
	}

	s.cac.Set(chunks[1], true, time.Hour*10)

	return nil
}

func (s *authService) CheckTokenValidity(token string) error {
	_, b := s.cac.Get(token)
	if b {
		return errorutils.ErrUnauthorized
	}

	return nil
}

func (s *authService) ReadUser(username string) (*model.User, error) {
	u, err := s.repository.ReadBy("username", username)
	if err != nil {
		return nil, errorutils.NewError(fiber.StatusBadRequest, err.Error())
	}

	return u, nil
}

func (s *authService) GetJWTKey() string {
	return s.conf.Security.Jwt.Key
}

func (s *authService) GetJWTExp() string {
	return s.conf.Security.Jwt.Exp
}
