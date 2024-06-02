package playlists

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := router.Group("/playlists")
    routes.POST("/", h.CreatePlaylist)
    routes.GET("/", h.GetPlaylists)
    routes.GET("/my", h.MyPlaylists)
    routes.PUT("/:id/add_song", h.AddSong)
    routes.GET("/:id", h.GetPlaylist)
    routes.PUT("/:id", h.UpdatePlaylist)
    routes.DELETE("/:id", h.DeletePlaylist)
}

