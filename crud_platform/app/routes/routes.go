package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	router *gin.Engine
}

func NewRoutes() Routes {
	r := Routes{
        router: gin.Default(),
    }

	return r
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run()
}