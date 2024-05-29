package songs

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

    routes := router.Group("/Songs")
    routes.POST("/", h.AddSong)
    routes.GET("/", h.GetSongs)
    routes.GET("/:id", h.GetSong)
    routes.PUT("/:id", h.UpdateSong)
    routes.DELETE("/:id", h.DeleteSong)
}

