package accesslog

import (
	"context"
	"time"

	"github.com/MedHubUz/access-log/internal/entity"
	"github.com/MedHubUz/access-log/pkg/rand"
)

type accessLogUseCase struct {
	contextTimeout time.Duration
	accessLogRepo  entity.AccessLogRepository
}

// new usecase ...
func NewUsecase(timeout time.Duration, accessLogRepo entity.AccessLogRepository) entity.AccessLogUsecase {
	return &accessLogUseCase{
		contextTimeout: timeout,
		accessLogRepo:  accessLogRepo,
	}
}

// before create
func (u *accessLogUseCase) beforeCreate(accessLog *entity.AccessLog) error {
	if len(accessLog.ID) == 0 {
		accessLog.ID = rand.String(20)
	}
	accessLog.CreatedAt = time.Now().UTC()
	return nil
}

// create
func (u *accessLogUseCase) Create(ctx context.Context, accessLog *entity.AccessLog) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	u.beforeCreate(accessLog)
	return u.accessLogRepo.Create(ctx, accessLog)
}

// read all
func (u *accessLogUseCase) ReadAll(ctx context.Context, parameters map[string][]string) ([]*entity.AccessLog, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.accessLogRepo.ReadAll(ctx, parameters)
}
