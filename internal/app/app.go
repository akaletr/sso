package app

import (
	grpcapp "cmd/sso/main.go/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, port)
	return &App{
		GRPCServer: grpcApp,
	}
}
