package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Author *Author
	AuthorID *uint
	Album *Album
	AlbumID *uint
	Playlists []*Playlist `gorm:"many2many:songs_playlists;"`
	Name string
	FileName string `gorm:"unique:true"`
}
