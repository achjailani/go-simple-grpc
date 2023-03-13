package persistence

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/domain/repository"
	"gorm.io/gorm"
)

type HttpLogRepo struct {
	db *gorm.DB
}

// NewHttpLogRepository is constructor
func NewHttpLogRepository(db *gorm.DB) *HttpLogRepo {
	return &HttpLogRepo{db: db}
}

var _ repository.HttpLogRepositoryInterface = &HttpLogRepo{}

func (h *HttpLogRepo) Save(ctx context.Context, log *entity.HttpLog) error {
	return h.db.WithContext(ctx).Create(log).Error
}

func (h *HttpLogRepo) Find(ctx context.Context, id int) (*entity.HttpLog, error) {
	var log entity.HttpLog

	err := h.db.WithContext(ctx).Where("id = ?", id).Take(&log).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (h *HttpLogRepo) Get(ctx context.Context) ([]*entity.HttpLog, error) {
	var logs []*entity.HttpLog
	err := h.db.WithContext(ctx).Find(logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}
