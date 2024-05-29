package songs

import (
    "net/http"

    "crud-platform/app/models"
    "github.com/gin-gonic/gin"
)

type UpdateSongRequestBody struct {
    Name      string `json:"title"`
    AuthorID  uint `json:"author_id"`
    AlbumID   uint `json:"album_id"`
    Year 		  uint `json:"year"`
}

func (h handler) UpdateSong(ctx *gin.Context) {
    id := ctx.Param("id")
    body := UpdateSongRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var song models.Song
	var author models.Author
	var album models.Album

    if result := h.DB.First(&song, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.FirstOrCreate(&author, body.AuthorID)
	h.DB.FirstOrCreate(&album, body.AlbumID)

	song.Name = body.Name
	song.Year = body.Year

	h.DB.Save(&song)

	h.DB.Model(&author).Association("Songs").Append(&song)
	h.DB.Model(&album).Association("Songs").Append(&song)

    ctx.JSON(http.StatusOK, &song)
}


