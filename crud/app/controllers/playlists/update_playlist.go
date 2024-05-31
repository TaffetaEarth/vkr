package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type UpdatePlaylistRequestBody struct {
    Name      string `json:"name"`
}

func (h handler) UpdatePlaylist(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdatePlaylistRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var playlist models.Playlist

    if result := h.DB.First(&playlist, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    playlist.Name = body.Name

    if result := h.DB.Save(&playlist); result.Error != nil {
        ctx.JSON(http.StatusBadRequest, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &playlist)
}


