package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateAlbumRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    SongsIDs  []uint `json:"song_ids"`
    Year      uint `json:"year"`
}

func (h handler) UpdateAlbum(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateAlbumRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var album models.Album
    
  if result := h.DB.First(&album, id); result.Error != nil {
      ctx.JSON(http.StatusNotFound, result.Error)
      return
  }
  album.Name = body.Name
  album.Year = body.Year

  if result := h.DB.Save(&album); result.Error != nil {
    ctx.JSON(http.StatusBadRequest, result.Error)
    return
  }    
  var author models.Author

	if body.AuthorID != 0 {
    h.DB.FirstOrCreate(&author, body.AuthorID)
    album.Author = &author
    h.DB.Model(&author).Association("Albums").Append(&album)
  } else {
      album.AuthorID = nil
  }
        
  var songsArray []models.Song
  h.DB.Find(&songsArray, body.SongsIDs)
	h.DB.Model(&album).Association("Songs").Replace(&songsArray)

  ctx.JSON(http.StatusOK, &album)
}


