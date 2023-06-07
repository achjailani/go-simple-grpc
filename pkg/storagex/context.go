package storagex

import "context"

// StorageX is a struct
type StorageX struct {
	strategy StorageStrategy
}

// NewStorageX is a constructor
func NewStorageX(strategy StorageStrategy) *StorageX {
	return &StorageX{strategy: strategy}
}

// Store is a method
func (s *StorageX) Store(ctx context.Context, object *FileUploadObject) error {
	return s.strategy.store(ctx, object)
}
