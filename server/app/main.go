package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/TaffetaEarth/vkr/minioclient"
	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

var fileName = "ccaed2bf-4ce7-4232-86c1-886437f8a614"

func main() {
	minioClient := minioclient.SetupMinioClient()
	r := gin.Default()

	// store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))

  // if err != nil {
  //   fmt.Println(err)
  // }
	// r.Use(sessions.Sessions("mysession", store))

	r.GET("/test_stream", func(c *gin.Context) {
		position := c.Query("position")
		// playerID := c.Query("playerId")
		// session := sessions.Default(c)

		var player Player
		var err error

		// if player, err = getCachedPlayer(session, playerID); err != nil {
		// 	player, err = getMinioPlayer(minioClient, fileName)
		// }

    player, err = getMinioPlayer(minioClient, fileName)

    fmt.Println("player created")

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

    fmt.Println("load chunks started")
    
    
    go player.startStream()
    fmt.Println("stream started")

		w := c.Writer
		header := w.Header()
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "audio/mpeg")

		w.WriteHeader(http.StatusOK)
		for {
			w.Write(<- player.stream)
			w.(http.Flusher).Flush()
			// select {
			// 	case <-c.Request.Context().Done():
			// 		// session.Set(playerID, player)
			// 		// session.Save()
			// 		return
			// }
		}
	})

	r.Run()
}

func (player *Player) loadChunks() {
	var err error

	for err != io.EOF {
		buffer := make([]byte, 4 * 44)

		bytesRead, err := player.file.Read(buffer)
	
		player.chunks = append(player.chunks, Chunk{content: buffer[:bytesRead], timestamp: i})

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				break
			}
		}
	}
}

func (player *Player) startStream() {
  for {
    for i := 0; i < len(player.chunks); i++ {
      if player.chunks[i].timestamp < player.currentStop {
        continue
      }
      player.stream <- player.chunks[i].content
    }
  }
}

func (player *Player) rewind() {
	chunks := player.chunks
	if chunks[0].timestamp > player.currentStop || chunks[len(chunks)-1].timestamp < player.currentStop {
		stats, _ := player.file.Stat()
		position := int64(player.currentStop) * stats.Size / int64(player.length)
		player.file.Seek(position, 0)
	}
}

func getCachedPlayer(session sessions.Session, playerID string) (Player, error) {
	cachedPlayer := session.Get(playerID)

	if cachedPlayer != nil {
		return cachedPlayer.(Player), nil
	}

	return Player{}, errors.New("no chached player for playerID")
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

type Player struct {
	file *minio.Object
	stream chan []byte
	length int64
	currentStop int64
}