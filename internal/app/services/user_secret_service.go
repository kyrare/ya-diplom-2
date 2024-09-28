package services

import (
	"context"

	"github.com/google/uuid"
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
	fileRepository repository.UserSecretFileRepository,
) interfaces.UserSecretService {
	return &UserSecretService{
		secretRepository: secretRepository,
		fileRepository:   fileRepository,
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

	createdSecret, err := s.secretRepository.Create(ctx, validatedSecret)
	if err != nil {
		return nil, err
	}

	fileData, err := (*command.SecretData).GetData()
	if err != nil {
		return nil, err
	}

	err = s.fileRepository.Store(ctx, createdSecret.Id, fileData)
	if err != nil {
		return nil, err
	}

	return createdSecret, nil
}

func (s *UserSecretService) GetAllForUser(ctx context.Context, userId uuid.UUID) ([]*entities.UserSecret, error) {
	secrets, err := s.secretRepository.GetAllForUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, secret := range secrets {
		d, err := s.fileRepository.Get(ctx, secret.Id)
		if err != nil {
			return nil, err
		}

		data, err := entities.MakeUserSecretData(secret.Type, d)
		if err != nil {
			return nil, err
		}
		secret.Data = &data
	}

	return secrets, nil
}
