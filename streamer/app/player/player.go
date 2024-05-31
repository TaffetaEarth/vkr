package player

import (
	"fmt"
	"io"
	"strconv"

	"github.com/minio/minio-go/v7"
)

type Player struct {
	File *minio.Object
	Stream chan []byte
	Length int64
	CurrentStop int64
	EndSignal chan struct{}
}

func (player *Player) LoadChunks() {
	defer player.File.Close()
	
	var err error

  stat, _ := player.File.Stat()
  position := player.CurrentStop * stat.Size / player.Length
  player.File.Seek(position, 0)

	for err != io.EOF {
		select {
		case <-player.EndSignal:
			return
		default:
			buffer := make([]byte, 10000)

			bytesRead, err := player.File.Read(buffer)
		
			player.Stream <- buffer[:bytesRead]

			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
					break
				}
			}
		} 
	}
}

func (p Player) Rewind(position string){
	intPosition, err := strconv.Atoi(position) 
	p.CurrentStop = int64(intPosition)

		if err != nil {
			fmt.Println(err)
		}
}
