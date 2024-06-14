package playlists

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdatePlaylistRequestBody struct {
    Name      string `json:"name"`
}

func (h handler) UpdatePlaylist(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdatePlaylistRequestBody{}

    currentUserId, _ := ctx.Get("currentUserId")
    userAdmin, _ := ctx.Get("userAdmin")

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var playlist models.Playlist

    var result *gorm.DB

    if userAdmin.(bool) {
        result = h.DB.First(&playlist, id)
    } else {
        result = h.DB.Where("user_id = ?", currentUserId).First(&playlist, id)
    }

    if result.Error != nil {
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


