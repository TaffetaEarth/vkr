package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	name string `gorm:"unique:true"`
	albums []Album
	songs []Song
}