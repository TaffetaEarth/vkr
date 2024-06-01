package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"crud/app/controllers/albums"
	"crud/app/controllers/authors"
	"crud/app/controllers/playlists"
	"crud/app/controllers/songs"
	"crud/app/controllers/users"
	"crud/app/grpc"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, cl grpc.Client) {
	songs.RegisterRoutes(router, db)
	authors.RegisterRoutes(router, db)
	albums.RegisterRoutes(router, db)
	playlists.RegisterRoutes(router, db)
	users.RegisterRoutes(router, &cl)
}