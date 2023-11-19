package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message any `json:"message"`
}

type ResponseWithData struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func Data(ctx *gin.Context, status int, response ResponseWithData) {
	ctx.JSON(status, response)
}

func InternalServerError(ctx *gin.Context) {
	Message(ctx, http.StatusInternalServerError, "Internal Server Error")
}
func Message(ctx *gin.Context, status int, message any) {
	ctx.JSON(status, ResponseMessage{Message: message})
}
