package main

import (
	// "log"
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/redis/go-redis/v9"

	"crud/app/controllers"
	"crud/app/db"
	"crud/app/grpc"
	"crud/app/middlewares/auth"
)


func main() {
  r := gin.Default()
	metricRouter := gin.Default()

	m := ginmetrics.GetMonitor()	
	// use metric middleware without expose metric path
	m.SetMetricPath("/metrics")
	// m.UseWithoutExposingEndpoint(r)
	// set metric path expose to metric router
	// m.Expose(metricRouter)
	m.Use(r)

  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
  grpcClient, _ := grpc.New(context.Background(), logger, "sso:44044", 10*time.Hour, 10)
	redisClient := initRedisClient()

  r.Use(auth.AuthChecker(logger, "secret", *grpcClient))
  r.Use(gin.Recovery())

  dbHandler := db.Init()
  controllers.RegisterRoutes(r, dbHandler, *grpcClient, redisClient)
	go func() {
		_ = metricRouter.Run(":8082")
	}()

  r.Run()	
}

func initRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "change-me",            
		DB:       0,              
	})
}
