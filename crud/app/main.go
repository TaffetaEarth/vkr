package main

import (
	// "log"
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"crud/app/controllers"
	"crud/app/db"
	"crud/app/grpc"
	"crud/app/middlewares/auth"
)


func main() {
  r := gin.Default()
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  grpcClient, _ := grpc.New(context.Background(), logger, "sso:44044", 10*time.Hour, 10)
	redisClient := initRedisClient()

  r.Use(auth.AuthChecker(logger, "secret", *grpcClient))
  r.Use(gin.Recovery())

  dbHandler := db.Init()
  controllers.RegisterRoutes(r, dbHandler, *grpcClient, redisClient)

  r.Run()	
}

func initRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "change-me",            
		DB:       0,              
	})
}
