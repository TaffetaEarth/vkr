package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Year uint
	Name string
	Author Author
	AuthorID *uint
	Songs []Song
}
