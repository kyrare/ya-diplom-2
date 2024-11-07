package interfaces

import (
	"context"

	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type AuthService interface {
	Login(ctx context.Context, productCommand *command.LoginCommand) (*command.LoginCommandResult, error)
	GetUserByToken(ctx context.Context, token string) (*entities.User, error)
	HashPassword(password string) (string, error)
	CheckUserPassword(user *entities.User, password string) bool
}
