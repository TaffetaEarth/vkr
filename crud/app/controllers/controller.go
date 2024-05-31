package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"crud/app/controllers/songs"
	"crud/app/controllers/authors"
	"crud/app/controllers/albums"
	"crud/app/controllers/playlists"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	songs.RegisterRoutes(router, db)
	authors.RegisterRoutes(router, db)
	albums.RegisterRoutes(router, db)
	playlists.RegisterRoutes(router, db)
}