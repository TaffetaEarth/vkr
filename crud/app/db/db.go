package db

import (
	"log"

	"github.com/TaffetaEarth/vkr/crud/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=db user=postgres password=change_me dbname=exam_db port=5432"

func Init() *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    checkErr(db.AutoMigrate(&models.Album{}))
  	checkErr(db.AutoMigrate(&models.Author{}))
  	checkErr(db.AutoMigrate(&models.Playlist{}))
  	checkErr(db.AutoMigrate(&models.Song{}))

    return db
}

func checkErr(err error) {
	if err != nil {
	  log.Fatalln(err)
	}
  }