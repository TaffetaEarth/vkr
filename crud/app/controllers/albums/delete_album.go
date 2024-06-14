package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteAlbum(ctx *gin.Context) {
    id := ctx.Param("id")

    userAdmin, _ := ctx.Get("userAdmin")

    var album models.Album

    if result := h.DB.First(&album, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    if userAdmin.(bool) {
        h.DB.First(&album, id)
    } else {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    h.DB.Delete(&album)

    ctx.Status(http.StatusOK)
}
