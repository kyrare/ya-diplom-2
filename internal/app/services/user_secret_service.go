package services

import (
	"context"

	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/domain/repository"
)

type UserSecretService struct {
	secretRepository repository.UserSecretRepository
	fileRepository   repository.UserSecretFileRepository
}

func NewUserSecretService(
	secretRepository repository.UserSecretRepository,
	// fileRepository repository.UserSecretFileRepository,
) interfaces.UserSecretService {
	return &UserSecretService{
		secretRepository: secretRepository,
		//fileRepository:   fileRepository,
	}
}

func (s *UserSecretService) Create(ctx context.Context, command *command.CreateUserSecretCommand) (*entities.UserSecret, error) {
	newSecret := entities.NewUserSecret(
		command.User,
		command.SecretType,
		command.SecretName,
		command.SecretData,
	)

	validatedSecret, err := entities.NewValidatedUserSecret(newSecret)
	if err != nil {
		return nil, err
	}

	createdSecret, err := s.secretRepository.Create(validatedSecret)

	if err != nil {
		return nil, err
	}

	// todo store file

	return createdSecret, nil
}
