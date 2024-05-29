package songs

import (
    "net/http"

    "crud-platform/app/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetSong(ctx *gin.Context) {
    id := ctx.Param("id")

    var song models.Song

    if result := h.DB.First(&song, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &song)
}
