package authors

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

    routes := router.Group("/authors")
    routes.POST("/", h.CreateAuthor)
    routes.GET("/", h.GetAuthors)
    routes.GET("/:id", h.GetAuthor)
    routes.PUT("/:id", h.UpdateAuthor)
    routes.DELETE("/:id", h.DeleteAuthor)
}

