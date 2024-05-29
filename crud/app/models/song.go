package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Author Author
	Album Album
	Year uint
	Name string
	FileName string `gorm:"unique:true"`
}

