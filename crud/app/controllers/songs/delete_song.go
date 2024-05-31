package songs

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteSong(ctx *gin.Context) {
    id := ctx.Param("id")

    var song models.Song

    if result := h.DB.First(&song, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&song)

    ctx.Status(http.StatusOK)
}
