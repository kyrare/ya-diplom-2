package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.ValidatedUser) (*entities.User, error)
	FindByIDs(ctx context.Context, id []uuid.UUID) ([]*entities.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByLogin(ctx context.Context, login string) (*entities.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
