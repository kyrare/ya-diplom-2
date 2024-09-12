package interfaces

import (
	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
)

type UserService interface {
	CreateUser(productCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	FindUserById(id uuid.UUID) (*command.FindUserByIdCommandResult, error)
	FindUserByLogin(id uuid.UUID) (*command.FindUserByLoginCommandResult, error)
	Delete(id uuid.UUID) error
}
