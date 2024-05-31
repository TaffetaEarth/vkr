package albums

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateSongRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    SongsIDs  []uint `json:"song_ids"`
    Year      uint `json:"year"`
}

func (h handler) UpdateAlbum(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateSongRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

	var author models.Author
	var album models.Album

    if result := h.DB.First(&album, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    
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

	h.DB.Model(&album).Association("Songs").Replace(&songsArray)

    ctx.JSON(http.StatusCreated, &album)
}


