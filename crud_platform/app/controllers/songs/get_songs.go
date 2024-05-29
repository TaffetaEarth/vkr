package songs

import (
    "net/http"

    "crud-platform/app/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetSongs(ctx *gin.Context) {
    var songs []models.Song

    if result := h.DB.Find(&songs); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &songs)
}
