package command

import (
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type CreateUserSecretCommand struct {
	User       *entities.User
	SecretType entities.UserSecretType
	SecretName string
	SecretData entities.UserSecretData
}
