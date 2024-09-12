package repository

import (
	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type UserRepository interface {
	Create(user *entities.ValidatedUser) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindByLogin(login string) (*entities.User, error)
	Delete(id uuid.UUID) error
}
