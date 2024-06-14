package authors

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateAuthorRequestBody struct {
    Name      string `json:"name"`
    AlbumIDs  []uint `json:"album_ids"`
    SongsIDs  []uint `json:"song_ids"`
}

func (h handler) UpdateAuthor(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateAuthorRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var author models.Author

    if result := h.DB.First(&author, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    author.Name = body.Name

    if result := h.DB.Save(&author); result.Error != nil {
        ctx.JSON(http.StatusBadRequest, result.Error)
        return
    }

    var songsArray []models.Song
    var albumsArray []models.Album

    h.DB.Find(&songsArray, body.SongsIDs)
    h.DB.Find(&albumsArray, body.AlbumIDs)

	h.DB.Model(&author).Association("Songs").Replace(&songsArray)
	h.DB.Model(&author).Association("Albums").Replace(&albumsArray)

    ctx.JSON(http.StatusCreated, &author)
}


