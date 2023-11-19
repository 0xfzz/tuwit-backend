package lib

import "github.com/gin-gonic/gin"

type Api struct {
	Route *gin.RouterGroup
}

func NewApi(app Router) Api {
	return Api{
		Route: app.Gin.Group("api"),
	}
}
