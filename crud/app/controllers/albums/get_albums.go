package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbums(ctx *gin.Context) {
    var albums []models.Album

    if result := h.DB.Find(&albums); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &albums)
}
