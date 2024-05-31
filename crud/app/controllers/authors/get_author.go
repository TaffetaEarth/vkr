package authors

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAuthor(ctx *gin.Context) {
    id := ctx.Param("id")

    var Author models.Author

    if result := h.DB.First(&Author, id); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &Author)
}
