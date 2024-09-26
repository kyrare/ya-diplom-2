package command

import (
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type ClientCreateUserSecretCommand struct {
	SecretType entities.UserSecretType
	SecretName string
	SecretData entities.UserSecretData
}
