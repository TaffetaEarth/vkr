package authors

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreateAuthorRequestBody struct {
    Name      string `json:"name"`
    AlbumIDs  uint `json:"album_ids"`
    SongsIDs  []uint `json:"song_ids"`
}

func (h handler) CreateAuthor(ctx *gin.Context) {
    body := CreateAuthorRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var author models.Author

    author.Name = body.Name

    if result := h.DB.Create(&author); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    var songsArray []models.Song
    var albumsArray []models.Album

    h.DB.Find(&songsArray, body.SongsIDs)
    h.DB.Find(&albumsArray, body.AlbumIDs)

	h.DB.Model(&author).Association("Songs").Append(&songsArray)
	h.DB.Model(&author).Association("Albums").Append(&albumsArray)

    ctx.JSON(http.StatusCreated, &author)
}

