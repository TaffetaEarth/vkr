package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) DeletePlaylist(ctx *gin.Context) {
    currentUserId, _ := ctx.Get("currentUserId")
    userAdmin, _ := ctx.Get("userAdmin")

	if currentUserId == 0 {
	    ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

    id := ctx.Param("id")

    var Playlist models.Playlist

    var result *gorm.DB

    if userAdmin.(bool) {
        result = h.DB.First(&Playlist, id)
    } else {
        result = h.DB.Where("user_id = ?", currentUserId).First(&Playlist, id)
    }

    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&Playlist)

    ctx.Status(http.StatusOK)
}
