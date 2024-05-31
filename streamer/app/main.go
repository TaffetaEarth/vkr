package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"streamer/app/minioclient"
	"streamer/app/player"
	"streamer/app/decoder"
	"streamer/app/statnotifier"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func main() {
	minioClient := minioclient.SetupMinioClient()
	r := gin.Default()

	r.GET("/stream/:file_name", func(c *gin.Context) {
		position := c.Query("position")
		fileName := c.Param("file_name")

		var player player.Player
		var err error

    player, err = getMinioPlayer(minioClient, fileName)

    if err != nil {
      fmt.Println(err)
    }

		sn := statnotifier.InitStatNotifier()
		sn.Publish(fileName)

		if position != "" { 
			player.Rewind(position)
		}

		go player.LoadChunks()

		w := c.Writer
		header := w.Header()
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "audio/mpeg")

		w.WriteHeader(http.StatusOK)
		for {
			w.Write(<- player.Stream)
			w.(http.Flusher).Flush()
			select {
			case <-c.Request.Context().Done():
				player.EndSignal <- struct{}{}
			default:
			}
		}
	})

	r.Run()
}

func getMinioPlayer(minioClient *minio.Client, fileName string) (player.Player, error) {
	ctx := context.Background()
	file, err := minioClient.GetObject(ctx, "music", fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return player.Player{}, errors.New("no minio player for file name")
	}

	return player.Player{Stream: make(chan []byte), File: file, Length: decoder.GetLength(file)}, nil
}