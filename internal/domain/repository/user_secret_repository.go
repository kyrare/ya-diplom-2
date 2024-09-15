package repository

import (
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserSecretRepository interface {
	Create(secret *entities.ValidatedUserSecret) (*entities.UserSecret, error)
}
