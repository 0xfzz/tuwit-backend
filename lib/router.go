package lib

import "github.com/gin-gonic/gin"

type Router struct {
	App *gin.Engine
}

func NewRouter() Router {
	app := gin.New()
	return Router{
		App: app,
	}
}
