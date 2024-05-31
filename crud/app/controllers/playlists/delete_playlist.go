package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeletePlaylist(ctx *gin.Context) {
    id := ctx.Param("id")

    var Playlist models.Playlist

    if result := h.DB.First(&Playlist, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&Playlist)

    ctx.Status(http.StatusOK)
}
