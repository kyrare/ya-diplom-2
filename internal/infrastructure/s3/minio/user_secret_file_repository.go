package minio

import (
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type UserSecretFileRepository struct {
	bucketName string
	client     *minio.Client
}

func NewMinioUserSecretFileRepository(bucketName string, client *minio.Client) *UserSecretFileRepository {
	return &UserSecretFileRepository{
		bucketName: bucketName,
		client:     client,
	}
}

func (r *UserSecretFileRepository) Store(objectId uuid.UUID, data []byte) error {
	//TODO implement me
	panic("implement me")
}
