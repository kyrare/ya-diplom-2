package repository

import "github.com/google/uuid"

type UserSecretFileRepository interface {
	Store(objectId uuid.UUID, data []byte) error
}
