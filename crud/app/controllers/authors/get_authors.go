package authors

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAuthors(ctx *gin.Context) {
    var Authors []models.Author

    if result := h.DB.Find(&Authors); result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &Authors)
}
