package users

import (
	"context"
	"crud/app/grpc"
	// "fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type grpcHandler struct {
    G *grpc.Client
}

type LoginRequestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func RegisterRoutes(router *gin.Engine, g *grpc.Client) {
    h := grpcHandler{
        G: g,
    }

    routes := router.Group("/user")
    routes.POST("/login", h.Login)
    routes.POST("/register", h.Register)
}

func (h grpcHandler) Login(ctx *gin.Context) {
    body := LoginRequestBody{}

    // currentUserId, _ := ctx.Get("currentUserId")

    // fmt.Println("current user is " + strconv.Itoa(int(currentUserId.(uint))))

    if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
    }
    token, err := h.G.Login(context.Background(), body.Email, body.Password)  
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h grpcHandler) Register(ctx *gin.Context) {
    body := LoginRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
    }
    token, err := h.G.Register(context.Background(), body.Email, body.Password) 
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
        return
    } 
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}
