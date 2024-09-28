package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
)

type UserService interface {
	Create(ctx context.Context, productCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	FindUserById(ctx context.Context, id uuid.UUID) (*command.FindUserByIdCommandResult, error)
	FindUserByLogin(ctx context.Context, id uuid.UUID) (*command.FindUserByLoginCommandResult, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
