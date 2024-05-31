package albums

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

    routes := router.Group("/albums")
    routes.POST("/", h.CreateAlbum)
    routes.GET("/", h.GetAlbums)
    routes.GET("/:id", h.GetAlbum)
    routes.PUT("/:id", h.UpdateAlbum)
    routes.DELETE("/:id", h.DeleteAlbum)
}

