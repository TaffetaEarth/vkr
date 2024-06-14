package authors

import (
	"fmt"
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

type CreateAuthorRequestBody struct {
    Name      string `json:"name"`
    AlbumsIDs []uint `json:"albums_ids"`
    SongsIDs  []uint `json:"songs_ids"`
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

    fmt.Println(body)

    h.DB.Where("id IN ?", body.SongsIDs).Find(&songsArray)
    h.DB.Where("id IN ?", body.AlbumsIDs).Find(&albumsArray)

	h.DB.Model(&author).Association("Songs").Append(&songsArray)
	h.DB.Model(&author).Association("Albums").Append(&albumsArray)

    ctx.JSON(http.StatusCreated, &author)
}

