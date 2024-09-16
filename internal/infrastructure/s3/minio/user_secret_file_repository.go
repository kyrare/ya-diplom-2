package minio

import (
	"bytes"
	"context"
	"fmt"

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

func (r *UserSecretFileRepository) Store(ctx context.Context, objectId uuid.UUID, data []byte) error {
	reader := bytes.NewReader(data)

	info, err := r.client.PutObject(
		ctx,
		r.bucketName,
		objectId.String(),
		reader,
		-1,
		minio.PutObjectOptions{
			//ContentType: "application/json",
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(info)

	return nil
}

func (r *UserSecretFileRepository) Get(ctx context.Context, objectId uuid.UUID) ([]byte, error) {
	object, err := r.client.GetObject(
		ctx,
		r.bucketName,
		objectId.String(),
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(object)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
