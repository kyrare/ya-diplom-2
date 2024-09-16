package repository

import (
	"context"

	"github.com/google/uuid"
)

type UserSecretFileRepository interface {
	Store(ctx context.Context, objectId uuid.UUID, data []byte) error
	Get(ctx context.Context, objectId uuid.UUID) ([]byte, error)
}
