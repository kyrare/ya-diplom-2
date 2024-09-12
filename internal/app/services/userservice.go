package services

import (
	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	newUser := entities.NewUser(
		userCommand.Login,
		userCommand.Password,
	)

	validatedUser, err := entities.NewValidatedUser(newUser)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.Create(validatedUser)
	if err != nil {
		return nil, err
	}

	return &command.CreateUserCommandResult{
		User: createdUser,
	}, nil
}

func (s *UserService) FindUserById(id uuid.UUID) (*command.FindUserByIdCommandResult, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return &command.FindUserByIdCommandResult{
		User: user,
	}, nil
}

func (s *UserService) FindUserByLogin(id uuid.UUID) (*command.FindUserByLoginCommandResult, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return &command.FindUserByLoginCommandResult{
		User: user,
	}, nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.userRepository.Delete(id)
}
