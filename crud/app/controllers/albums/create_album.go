package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreateAlbumRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    SongsIDs  []uint `json:"song_ids"`
    Year      uint `json:"year"`
}

func (h handler) CreateAlbum(ctx *gin.Context) {
    body := CreateAlbumRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var author models.Author
	var album models.Album

	h.DB.FirstOrCreate(&author, body.AuthorID)

    album.Name = body.Name
    album.Year = body.Year
    album.Author = author

    if result := h.DB.Create(&album); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    var songsArray []models.Song

    h.DB.Find(&songsArray, body.SongsIDs)

	h.DB.Model(&album).Association("Songs").Append(&songsArray)

    ctx.JSON(http.StatusCreated, &album)
}

