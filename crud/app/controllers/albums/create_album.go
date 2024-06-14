package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreateAlbumRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    SongsIDs  []uint
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

	if body.AuthorID != 0 {
    h.DB.FirstOrCreate(&author, body.AuthorID)
    album.Author = &author
    h.DB.Model(&author).Association("Albums").Append(&album)
  } else {
    album.AuthorID = nil
  }

    album.Name = body.Name
    album.Year = body.Year

    if result := h.DB.Create(&album); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    var songsArray []models.Song

    if len(body.SongsIDs) == 0 {
        body.SongsIDs = []uint{0}
    }
    h.DB.Find(&songsArray, body.SongsIDs)

	h.DB.Model(&album).Association("Songs").Append(&songsArray)

    ctx.JSON(http.StatusCreated, &album)
}

