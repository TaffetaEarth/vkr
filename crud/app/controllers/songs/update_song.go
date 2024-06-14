package songs

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateSongRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint   `json:"author_id"`
    AlbumID   uint   `json:"album_id"`
    Year 	  uint   `json:"year"`
}

func (h handler) UpdateSong(ctx *gin.Context) {
  id := ctx.Param("id")
  body := UpdateSongRequestBody{}

  if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
  }

  var song models.Song
	var author models.Author
	var album models.Album

  if result := h.DB.First(&song, id); result.Error != nil {
      ctx.JSON(http.StatusNotFound, result.Error)
      return
  }

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

	h.DB.Save(&song)

	h.DB.Model(&song).Association("Authors").Replace(&author)
	h.DB.Model(&song).Association("Album").Replace(&album)

  ctx.JSON(http.StatusOK, &song)
}


