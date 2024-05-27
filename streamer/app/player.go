package main

import (
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type Player struct {
	file *minio.Object
	stream chan []byte
	length int64
	currentStop int64
}

func (player *Player) loadChunks() {
	var err error

  stat, _ := player.file.Stat()
  position := player.currentStop * stat.Size / player.length
  player.file.Seek(position, 0)

	for err != io.EOF {
		buffer := make([]byte, 10000)

		bytesRead, err := player.file.Read(buffer)
	
		player.stream <- buffer[:bytesRead]

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				break
			}
		}
	}
}
