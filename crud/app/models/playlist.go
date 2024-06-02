package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	UserID *uint
	Songs []*Song `gorm:"many2many:songs_playlits;"`
	Name string
}
