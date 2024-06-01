package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/cmd/sso/internal/config"
	"sso/internal/app"
	"syscall"
)

func main() {  
	cfg := config.MustLoad()

	log := setupLogger()

	application := app.New(log, cfg.GRPC.Port, cfg.TokenTTL)

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Waiting for SIGINT (pkill -2) or SIGTERM
	<-stop

	// initiate graceful shutdown
	application.GRPCServer.Stop() // Assuming GRPCServer has Stop() method for graceful shutdown
	log.Info("Gracefully stopped")
}

func setupLogger() *slog.Logger {
	return slog.New(
					slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
				)
}
