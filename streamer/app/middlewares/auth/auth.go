package auth

import (
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


func AuthChecker(log *slog.Logger, appSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := extractBearerToken(c.GetHeader("Authorization"))
		if tokenStr == "" {
			fmt.Println("no token provided")
			c.Set("userAuthorized", false)
			return
		}
		claims := jwt.MapClaims{}
		
		_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
		}

			return []byte(appSecret), nil
		})
	
		if err != nil {
				log.Warn("failed to parse token")
				// But if token is invalid, we shouldn't handle request
				c.Set("userAuthorized", false)
				return
		}

		c.Set("userAuthorized", true)
	}
}
