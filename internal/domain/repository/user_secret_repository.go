package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserSecretRepository interface {
	Create(ctx context.Context, secret *entities.ValidatedUserSecret) (*entities.UserSecret, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAllForUser(ctx context.Context, id uuid.UUID) ([]*entities.UserSecret, error)
}
