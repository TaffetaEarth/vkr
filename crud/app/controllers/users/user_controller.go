package users

import (
	"context"
	"crud/app/grpc"
	"net/http"

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

    if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
    }
    token, _ := h.G.Login(context.Background(), body.Email, body.Password)  
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h grpcHandler) Register(ctx *gin.Context) {
    body := LoginRequestBody{}

    if err := ctx.BindJSON(&body); err != nil {
      ctx.JSON(http.StatusBadRequest, err)
      return
    }
    token, _ := h.G.Register(context.Background(), body.Email, body.Password)  
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}
