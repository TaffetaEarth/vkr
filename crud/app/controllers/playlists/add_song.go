package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type AddSongToPlaylistRequestBody struct {
	ID      uint `json:"id"`
}

func (h handler) AddSong(ctx *gin.Context) {
		currentUserId, _ := ctx.Get("currentUserId")
		id := ctx.Param("id")

		if currentUserId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}

    var playlist models.Playlist

    if result := h.DB.Where("user_id = ?", currentUserId).First(&playlist, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

		body := AddSongToPlaylistRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

		var song models.Song

    if result := h.DB.First(&song, body.ID); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

		h.DB.Model(&playlist).Association("Songs").Append(&song)

    ctx.JSON(http.StatusOK, &playlist)
}
