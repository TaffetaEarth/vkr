package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	UserID uint
	songs []Song
	name string
}

