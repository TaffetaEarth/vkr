package songs

import (
	"context"
	"net/http"

	"crud/app/minioclient"
	"crud/app/models"
	"crud/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (h handler) UploadFile(ctx *gin.Context) {
	id := ctx.Param("id")

	var song models.Song
	if result := h.DB.First(&song, id); result.Error != nil {
			ctx.JSON(http.StatusNotFound, result.Error)
			return
	}

  file, err := ctx.FormFile("file")
  utils.CheckErr(err)

  openFile, err := file.Open()
  utils.CheckErr(err)

  if err == nil {
      c := context.Background()
      minioClient := minioclient.SetupMinioClient()
      name := uuid.NewString()
      _, err = minioClient.PutObject(c, "music", name, openFile, file.Size, minio.PutObjectOptions{ContentType: "audio/mpeg"})
      if err == nil {
          song.FileUrl = "localhost:8000/stream/" + name 
      }
  } else {
    ctx.JSON(http.StatusBadRequest, err.Error())
  }

	h.DB.Save(&song)

	ctx.JSON(http.StatusOK, &song)
}

