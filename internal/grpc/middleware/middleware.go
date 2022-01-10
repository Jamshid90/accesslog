package middleware

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Middleware struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Middleware {
	return &Middleware{logger}
}

func (m *Middleware) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// md, ok := metadata.FromIncomingContext(ctx)
	// if ok {
	// 	fmt.Println(md)
	// }

	h, err := handler(ctx, req)

	if err != nil {
		m.logger.Error("gRPC request",
			zap.String("method", info.FullMethod),
			zap.Duration("time", time.Since(start)))
		return h, err
	}

	m.logger.Info("gRPC request",
		zap.String("method", info.FullMethod),
		zap.Duration("time", time.Since(start)),
	)
	return h, err
}
