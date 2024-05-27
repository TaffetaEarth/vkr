package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TaffetaEarth/vkr/minioclient"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

var fileName = "ccaed2bf-4ce7-4232-86c1-886437f8a614"

func main() {
	minioClient := minioclient.SetupMinioClient()
	r := gin.Default()

	r.GET("/test_stream", func(c *gin.Context) {
		position := c.Query("position")

		var player Player
		var err error

    player, err = getMinioPlayer(minioClient, fileName)

    if err != nil {
      fmt.Println(err)
    }

	if position != "" {
		intPosition, err := strconv.Atoi(position) 
		player.currentStop = int64(intPosition)

      if err != nil {
        fmt.Println(err)
      }
	}

		go player.loadChunks()

		w := c.Writer
		header := w.Header()
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "audio/mpeg")

		w.WriteHeader(http.StatusOK)
		for {
			w.Write(<- player.stream)
			w.(http.Flusher).Flush()
		}
	})

	r.Run()
}

func getMinioPlayer(minioClient *minio.Client, fileName string) (Player, error) {
	ctx := context.Background()
	file, err := minioClient.GetObject(ctx, "music", fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return Player{}, errors.New("no minio player for file name")
	}

	return Player{stream: make(chan []byte), file: file, length: getLength(file)}, nil
}