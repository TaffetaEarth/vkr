package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	year uint
	name string
	author Author
	songs []Song
}


