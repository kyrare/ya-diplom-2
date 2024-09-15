package minio

import (
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(config *services.Config, logger *services.Logger) (*minio.Client, error) {
	conf := &config.Minio

	logger.Infof("crete new minio client, host: %s key: %s useSsl: %t", conf.Endpoint, conf.AccessKey, conf.UseSSL)

	client, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
