package songs

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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


  if body.AuthorID != 0 {
	  h.DB.FirstOrCreate(&author, body.AuthorID)
    song.Author = &author
    h.DB.Model(&author).Association("Songs").Append(&song)
  } else {
    song.AuthorID = nil
  }
  if body.AlbumID != 0 {
    h.DB.FirstOrCreate(&author, body.AuthorID)
    album.Author = &author
    song.Album = &album
    h.DB.Model(&album).Association("Songs").Append(&song)
  } else {
    song.AlbumID = nil
  }

  song.Name = body.Name
  song.FileName = uuid.NewString()

  if result := h.DB.Create(&song); result.Error != nil {
      ctx.JSON(http.StatusNotFound, result.Error)
      return
  }

  ctx.JSON(http.StatusCreated, &song)
}

