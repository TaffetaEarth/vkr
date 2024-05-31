package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteAlbum(ctx *gin.Context) {
    id := ctx.Param("id")

    var album models.Album

    if result := h.DB.First(&album, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&album)

    ctx.Status(http.StatusOK)
}
