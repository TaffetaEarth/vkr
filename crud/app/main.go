package main

import (
	// "log"
	"github.com/gin-gonic/gin"

	"crud/app/controllers"
	"crud/app/db"
)


func main() {
  r := gin.Default()
  r.Use(gin.Recovery())

  dbHandler := db.Init()
  controllers.RegisterRoutes(r, dbHandler)

  r.Run()	
}
