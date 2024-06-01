package main

import (
	// "log"
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"crud/app/controllers"
	"crud/app/db"
	"crud/app/grpc"
	"crud/app/middlewares/auth"
)


func main() {
  r := gin.Default()
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  grpcClient, _ := grpc.New(context.Background(), logger, "sso:44044", 10*time.Hour, 10)

  r.Use(auth.AuthChecker(logger, "secret", *grpcClient))
  r.Use(gin.Recovery())

  dbHandler := db.Init()
  controllers.RegisterRoutes(r, dbHandler, *grpcClient)

  r.Run()	
}
