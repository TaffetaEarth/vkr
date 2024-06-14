package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbum(ctx *gin.Context) {
    id := ctx.Param("id")

    var album models.Album

    if result := h.DB.Preload("Songs").Preload("Author").First(&album, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &album)
}
