package app

import (
    "log/slog"
    "time"

    "sso/internal/app/grpc"
    "sso/internal/services/auth"
    "sso/internal/storage/postgres"
)

type App struct {
    GRPCServer *grpc.App
    Storage postgres.DbHandler
}

func New(
    log *slog.Logger,
    grpcPort int,
    tokenTTL time.Duration,
) *App {
    storage := postgres.Init()
    
    authService := auth.New(log, storage, storage, tokenTTL)

    grpcApp := grpc.New(log, authService, grpcPort)

    return &App{
        GRPCServer: grpcApp,
        Storage: *storage,
    }
}
