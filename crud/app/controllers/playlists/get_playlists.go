package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPlaylists(ctx *gin.Context) {
    var Playlists []models.Playlist

    if result := h.DB.Find(&Playlists); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &Playlists)
}
