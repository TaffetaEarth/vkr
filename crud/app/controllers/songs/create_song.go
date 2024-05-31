package songs

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreateSongRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    AlbumID   uint `json:"album_id"`
    Year uint `json:"year"`
}

func (h handler) CreateSong(ctx *gin.Context) {
  body := CreateSongRequestBody{}

  if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
  }

  var song models.Song
  var author models.Author
  var album models.Album
  
	h.DB.FirstOrCreate(&author, body.AuthorID)
	h.DB.FirstOrCreate(&album, body.AlbumID)

    song.Name = body.Name
    song.Author = author
    album.Author = author
    song.Album = album

  if result := h.DB.Create(&song); result.Error != nil {
      ctx.JSON(http.StatusNotFound, result.Error)
      return
  }

	h.DB.Model(&author).Association("Songs").Append(&song)
	h.DB.Model(&album).Association("Songs").Append(&song)

  ctx.JSON(http.StatusCreated, &song)
}

