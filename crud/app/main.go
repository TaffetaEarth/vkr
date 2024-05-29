package main

import (
	// "log"
	"github.com/gin-gonic/gin"

	"github.com/TaffetaEarth/vkr/crud/app/controllers/songs"
	"github.com/TaffetaEarth/vkr/crud/app/db"
)


func main() {
  r := gin.Default()
  r.Use(gin.Recovery())

  dbHandler := db.Init()
  songs.RegisterRoutes(r, dbHandler)

	r.Run()
}

// func checkErr(err error) {
//   if err != nil {
//     log.Fatalln(err)
//   }
// }