package main

import (
	"log"
	"time"

	accesslogpb "github.com/MedHubUz/access-log/genproto/accesslog"
	"github.com/MedHubUz/access-log/internal/accesslog"
	configpkg "github.com/MedHubUz/access-log/internal/config"
	"github.com/MedHubUz/access-log/internal/database"
	grpcmiddleware "github.com/MedHubUz/access-log/internal/grpc/middleware"
	grpcserver "github.com/MedHubUz/access-log/internal/grpc/server"
	loggerpkg "github.com/MedHubUz/access-log/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	var (
		contextTimeout time.Duration
	)

	// initialization config
	config := configpkg.New()

	// initialization logger
	logger, err := loggerpkg.New(config.LogLevel, config.Environment)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	db, err := database.New(config)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// context timeout initialization
	contextTimeout, err = time.ParseDuration(config.Context.Timeout)
	if err != nil {
		log.Fatalf("Error during parse duration for context timeout : %v\n", err)
	}

	// repositories initialization
	accesslogRepo := accesslog.NewRepository(db)
	accesslogUsecase := accesslog.NewUsecase(contextTimeout, accesslogRepo)

	// gRPC server initialization
	middleware := grpcmiddleware.New(logger)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryInterceptor),
	)

	accesslogpb.RegisterAccessLogServiceServer(grpcServer, accesslog.NewgRPC(logger, accesslogUsecase))

	logger.Info("Listen gRPC", zap.String("url", config.RPCPort))
	if err := grpcserver.Run(config, grpcServer); err != nil {
		log.Fatalf("gRPC fatal to serve grpc server over %s %v", config.RPCPort, err)
	}
}
