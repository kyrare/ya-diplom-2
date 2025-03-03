package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repository.UserRepository
	tokenSecret    string
	tokenDuration  time.Duration
	logger         *Logger
}

type AuthJwtClaims struct {
	jwt.RegisteredClaims
	UID   uuid.UUID `json:"uid"`
	Login string    `json:"login"`
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func NewAuthService(
	userRepository repository.UserRepository,
	tokenSecret string,
	tokenDuration time.Duration,
	logger *Logger,
) interfaces.AuthService {
	return &AuthService{
		userRepository: userRepository,
		tokenSecret:    tokenSecret,
		tokenDuration:  tokenDuration,
		logger:         logger,
	}
}

func (s AuthService) Login(ctx context.Context, c *command.LoginCommand) (*command.LoginCommandResult, error) {
	if c.Login == "" || c.Password == "" {
		s.logger.Debug("Авторизации с пустым логином или паролем")
		return nil, ErrInvalidCredentials
	}

	user, err := s.userRepository.FindByLogin(ctx, c.Login)
	if err != nil {
		s.logger.Debugf("Авториазации завершилась с ошибкой при поиске пользователя, ошибка: %s", err.Error())
		return nil, err
	}
	if user == nil {
		s.logger.Debugf("При попытке авториазации пользователь с логином %s небыл надйен", c.Login)
		return nil, ErrInvalidCredentials
	}

	if !s.CheckUserPassword(user, c.Password) {
		s.logger.Debugf("При попытке авториазации с логином %s пароль не прошел проверку", c.Login)
		return nil, ErrInvalidCredentials
	}

	token, err := s.newJwtToken(user)
	if err != nil {
		s.logger.Error("Не удалось сгенерировать jwt токен, ошибка: %s", err.Error())
		return nil, err
	}

	return &command.LoginCommandResult{
		JwtToken: token,
	}, nil
}

func (s AuthService) GetUserByToken(ctx context.Context, t string) (*entities.User, error) {
	claims, err := s.parseJwtToken(t)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindById(ctx, claims.UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s AuthService) newJwtToken(user *entities.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// todo спросить про эту запись
	claims := token.Claims.(jwt.MapClaims)

	claims["uid"] = user.Id
	claims["login"] = user.Login
	claims["exp"] = time.Now().Add(s.tokenDuration).Unix()

	tokenString, err := token.SignedString([]byte(s.tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s AuthService) parseJwtToken(t string) (*AuthJwtClaims, error) {
	if t == "" {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(t, &AuthJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*AuthJwtClaims), nil
}

func (s AuthService) HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func (s AuthService) CheckUserPassword(user *entities.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
