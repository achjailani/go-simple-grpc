package storagex

import (
	"context"
)

// StorageStrategy is a contract
type StorageStrategy interface {
	store(ctx context.Context, file *FileUploadObject) error
}
