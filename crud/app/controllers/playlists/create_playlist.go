package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistRequestBody struct {
    Name      string `json:"name"`
    SongsIDs  []uint `json:"song_ids"`
}

func (h handler) CreatePlaylist(ctx *gin.Context) {
    body := CreatePlaylistRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var playlist models.Playlist
    playlist.Name = body.Name

    if result := h.DB.Create(&playlist); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    var songsArray []models.Song

    h.DB.Find(&songsArray, body.SongsIDs)
	h.DB.Model(&playlist).Association("Songs").Append(&songsArray)
    
    ctx.JSON(http.StatusCreated, &playlist)
}

