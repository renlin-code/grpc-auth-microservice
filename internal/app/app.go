package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/renlin-code/grpc-sso-microservice/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func NewApp(log *slog.Logger, port int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcApp.NewApp(log, port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
