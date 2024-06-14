package authors

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteAuthor(ctx *gin.Context) {
    id := ctx.Param("id")

    currentUserId, _ := ctx.Get("currentUserId")
    userAdmin, _ := ctx.Get("userAdmin")

    if currentUserId == 0 {
	    ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

    var Author models.Author

    if userAdmin.(bool) {
        h.DB.First(&Author, id)
    } else {
        ctx.JSON(http.StatusForbidden, gin.H{"status": 403, "error": "forbidden"})
        return
    }

    h.DB.Delete(&Author)

    ctx.Status(http.StatusOK)
}
