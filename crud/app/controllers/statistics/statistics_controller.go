package statistics

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"crud/app/models"
)

type redisHandler struct {
    R *redis.Client
    DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, rc *redis.Client, db *gorm.DB) {
    rh := redisHandler{
        R: rc,
        DB: db,
    }

    routes := router.Group("/charts")
    routes.GET("/", rh.GetCharts)
}

func (rh *redisHandler) GetCharts(ctx *gin.Context) {
    context := context.Background()
    result:= rh.R.ZRangeByScore(context, "statistics", &redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: 0, Count: 10}) 
    
    fileList, err := result.Result()
    
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
        return
    }
    if len(fileList) == 0{
        ctx.JSON(http.StatusNotFound, gin.H{"error": "no chart data found"})
        return
    }
    var songs []models.Song

    rh.DB.Raw(`select s.*
    from songs s
      join unnest(array[?]) with ordinality as x (file_name, ordering) 
        on s.file_name = x.file_name
    order by x.ordering
    `, strings.Join(fileList, ", ")).Scan(&songs)
    

    ctx.JSON(http.StatusOK, &songs)
}
