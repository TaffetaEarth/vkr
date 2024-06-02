package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistRequestBody struct {
    Name      string `json:"name"`
    SongsIDs  []uint 
}

func (h handler) CreatePlaylist(ctx *gin.Context) {
    currentUserIdString, _ := ctx.Get("currentUserId")

    currentUserId := currentUserIdString.(uint)

    if currentUserId == 0 {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
    }

    body := CreatePlaylistRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var playlist models.Playlist
    playlist.Name = body.Name
    playlist.UserID = &currentUserId

    if result := h.DB.Create(&playlist); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    var songsArray []models.Song

    h.DB.Find(&songsArray, body.SongsIDs)
	h.DB.Model(&playlist).Association("Songs").Append(&songsArray)
    
    ctx.JSON(http.StatusCreated, &playlist)
}

