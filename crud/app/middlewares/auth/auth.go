// internal/http-server/middleware/auth/auth.go

package auth

import (
	"crud/app/grpc"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func extractBearerToken(authHeader string) string {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
			return ""
	}

	return splitToken[1]
}


func AuthChecker(log *slog.Logger, appSecret string, grpcClient grpc.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := extractBearerToken(c.GetHeader("Authorization"))
		if tokenStr == "" {
			c.Set("current_user", nil)
			fmt.Println("no token provided")
			return
		}
		claims := jwt.MapClaims{}
		
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
		}

			return []byte(appSecret), nil
		})
	
		if err != nil {
				log.Warn("failed to parse token")
				// But if token is invalid, we shouldn't handle request
				c.Set("current_user", nil)
				return
		}

		log.Info("user authorized", slog.Any("claims", claims))

		userId := token.Claims.(jwt.MapClaims)["uid"].(uint)

		// Отправляем запрос для проверки, является ли пользователь админов
		isAdmin, err := grpcClient.IsAdmin(c, userId)
		if err != nil {
				log.Error("failed to check if user is admin")
				c.Set("user_admin", false)
				return
		}

		c.Set("user_admin", isAdmin)
		c.Set("current_user", userId)
	}
}
