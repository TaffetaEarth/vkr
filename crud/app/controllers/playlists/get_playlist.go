package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPlaylist(ctx *gin.Context) {
    id := ctx.Param("id")

    var Playlist models.Playlist

    if result := h.DB.Preload("Songs").First(&Playlist, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &Playlist)
}
