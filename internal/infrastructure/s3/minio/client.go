package minio

import (
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(endpoint, accessKey, secretKey string, useSSL bool, logger *services.Logger) (*minio.Client, error) {
	logger.Infof("crete new minio client, host: %s key: %s useSsl: %t", endpoint, accessKey, useSSL)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
