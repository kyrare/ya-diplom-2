package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
	authService    interfaces.AuthService
	logger         *Logger
}

func NewUserService(userRepository repository.UserRepository, authService interfaces.AuthService, logger *Logger) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
		authService:    authService,
		logger:         logger,
	}
}

func (s *UserService) Create(ctx context.Context, userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	newUser := entities.NewUser(
		userCommand.Login,
		userCommand.Password,
	)

	validatedUser, err := entities.NewValidatedUser(newUser)
	if err != nil {
		return nil, err
	}

	existUser, err := s.userRepository.FindByLogin(ctx, userCommand.Login)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, errors.New("такой пользователь уже существует")
	}

	hash, err := s.authService.HashPassword(validatedUser.Password)
	if err != nil {
		s.logger.Errorf("Не удалось сгенерировать хеш, ошибка: %s", err.Error())
		return nil, err
	}
	validatedUser.Password = hash

	createdUser, err := s.userRepository.Create(ctx, validatedUser)
	if err != nil {
		return nil, err
	}

	return &command.CreateUserCommandResult{
		User: createdUser,
	}, nil
}

func (s *UserService) FindUserById(ctx context.Context, id uuid.UUID) (*command.FindUserByIdCommandResult, error) {
	user, err := s.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &command.FindUserByIdCommandResult{
		User: user,
	}, nil
}

func (s *UserService) FindUserByLogin(ctx context.Context, id uuid.UUID) (*command.FindUserByLoginCommandResult, error) {
	user, err := s.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &command.FindUserByLoginCommandResult{
		User: user,
	}, nil
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.userRepository.Delete(ctx, id)
}
