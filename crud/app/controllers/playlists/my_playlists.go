package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) MyPlaylists(ctx *gin.Context) {
		currentUserId, _ := ctx.Get("currentUserId")

		if currentUserId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

    var Playlists []models.Playlist

    if result := h.DB.Preload("Songs").Where("user_id = ?", currentUserId).Find(&Playlists); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &Playlists)
}
