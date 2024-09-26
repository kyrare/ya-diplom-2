package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type ClientService interface {
	Login(ctx context.Context, command *command.LoginCommand) error
	Register(ctx context.Context, command *command.RegisterCommand) error
	GetUserSecrets(ctx context.Context) ([]*entities.UserSecret, error)
	GetUserSecretById(ctx context.Context, id uuid.UUID) (*entities.UserSecret, error)
	CreateUserSecret(ctx context.Context, command *command.ClientCreateUserSecretCommand) error
}
