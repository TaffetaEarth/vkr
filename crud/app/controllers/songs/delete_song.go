package songs

import (
	"net/http"

	"crud/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) DeleteSong(ctx *gin.Context) {
    id := ctx.Param("id")

    currentUserId, _ := ctx.Get("currentUserId")
    userAdmin, _ := ctx.Get("userAdmin")

	if currentUserId == 0 {
	    ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

    var song models.Song

    var result *gorm.DB

    if userAdmin.(bool) {
        result = h.DB.First(&song, id)
    } else {
        ctx.JSON(http.StatusForbidden, gin.H{"status": 403, "error": "forbidden"})
        return
    }

    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&song)

    ctx.Status(http.StatusOK)
}
