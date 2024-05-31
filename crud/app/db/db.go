package db

import (
	"crud/app/models"
	"crud/app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=postgres user=postgres password=change-me dbname=exam-db port=5432"

func Init() *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    utils.CheckErr(err)

    utils.CheckErr(db.AutoMigrate(&models.Album{}))
  	utils.CheckErr(db.AutoMigrate(&models.Author{}))
  	utils.CheckErr(db.AutoMigrate(&models.Playlist{}))
  	utils.CheckErr(db.AutoMigrate(&models.Song{}))

    return db
}