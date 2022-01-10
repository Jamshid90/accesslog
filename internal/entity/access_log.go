package entity

import (
	"context"
	"time"
)

type AccessLog struct {
	ID        string
	UserID    string
	Method    string
	URL       string
	Data      []byte
	CreatedAt time.Time
}

type AccessLogRepository interface {
	Create(ctx context.Context, log *AccessLog) error
	ReadAll(ctx context.Context, parameters map[string][]string) ([]*AccessLog, error)
}

type AccessLogUsecase interface {
	Create(ctx context.Context, log *AccessLog) error
	ReadAll(ctx context.Context, parameters map[string][]string) ([]*AccessLog, error)
}
