package storagex

import (
	"context"
	"github.com/minio/minio-go"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/utils"
)

// Minio is a struct
type Minio struct {
	client *minio.Client
	bucket string
}

// Store is a method
func (m *Minio) store(ctx context.Context, object *FileUploadObject) error {
	fileType, _ := utils.GetContentTypeFromFile(object.File)

	_, err := m.client.PutObjectWithContext(ctx, m.bucket, object.FileName, object.File, -1, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		return err
	}

	return nil
}

// NewMinio is a constructor
func NewMinio(cfg *config.Config) (*Minio, error) {
	client, err := minio.New(
		cfg.StorageConfig.StorageEndpoint,
		cfg.StorageConfig.StorageAccessKeyID,
		cfg.StorageConfig.StorageAccessKeySecret,
		false,
	)
	if err != nil {
		return nil, err
	}

	return &Minio{
		client: client,
		bucket: cfg.StorageConfig.StorageBucketName,
	}, nil
}

var _ StorageStrategy = (*Minio)(nil)
