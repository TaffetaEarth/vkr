package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `gorm:"unique:true"`
	Albums []Album
	Songs []Song
}
