package storagex

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github/achjailani/go-simple-grpc/config"
	"io"
	"time"
)

// GCS is a struct
type GCS struct {
	client    *storage.Client
	bucket    string
	path      string
	projectID string
}

// Store is a method
func (g *GCS) store(ctx context.Context, object *FileUploadObject) error {
	defer object.File.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := g.client.Bucket(g.bucket).Object(object.FileName)
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err := io.Copy(wc, object.File); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}

	return nil
}

// NewCGS is a constructor
func NewCGS(cfg *config.Config) (*GCS, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &GCS{
		client: client,
	}, nil
}

// requires to implement methods
var _ StorageStrategy = (*GCS)(nil)
