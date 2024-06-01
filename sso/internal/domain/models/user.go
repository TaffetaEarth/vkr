package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email    string
    PassHash []byte
    IsAdmin bool
}