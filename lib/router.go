package lib

import "github.com/gin-gonic/gin"

type Router struct {
	Gin *gin.Engine
}

func NewRouter() Router {
	app := gin.New()
	return Router{
		Gin: app,
	}
}
