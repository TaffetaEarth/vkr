package jwt

import (
	"time"

	"sso/internal/domain/models"

	"github.com/golang-jwt/jwt/v5"
)

var secret = "secret"

func NewToken(user models.User, duration time.Duration) (string, error) {  
	token := jwt.New(jwt.SigningMethodHS256)  

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)  
	claims["uid"] = user.ID  
	claims["email"] = user.Email  
	claims["exp"] = time.Now().Add(duration).Unix()

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(secret))  
	if err != nil {  
		 return "", err  
	}  

	return tokenString, nil  
}
