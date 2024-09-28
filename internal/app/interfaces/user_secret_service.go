package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserSecretService interface {
	Create(ctx context.Context, command *command.CreateUserSecretCommand) (*entities.UserSecret, error)
	GetAllForUser(ctx context.Context, userId uuid.UUID) ([]*entities.UserSecret, error)
}
