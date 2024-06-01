package postgres

import (
	"context"
	"fmt"
	"sso/internal/domain/models"
	"sso/internal/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=postgres user=postgres password=change-me dbname=exam-db port=5432"

type DbHandler struct {
	DB *gorm.DB
}

func Init() *DbHandler {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

    db.AutoMigrate(&models.User{})

    return &DbHandler{DB: db}
}

func (h *DbHandler) SaveUser(cxt context.Context, email string, passHash []byte) (uint, error) {
	const op = "storage.postgres.SaveUser"

	var user models.User

	user.Email = email
	user.PassHash = passHash

	if result := h.DB.Create(&user); result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	return user.ID, nil
}

func (h *DbHandler) User(cxt context.Context, email string) (models.User, error) {
	const op = "storage.postgres.User"

	var user models.User

	user.Email = email

	if result := h.DB.First(&user); result.Error != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return user, nil
}
