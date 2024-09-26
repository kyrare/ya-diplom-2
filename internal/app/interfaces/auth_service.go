package interfaces

import (
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type AuthService interface {
	Login(productCommand *command.LoginCommand) (*command.LoginCommandResult, error)
	HashPassword(password string) (string, error)
	CheckUserPassword(user *entities.User, password string) bool
}
