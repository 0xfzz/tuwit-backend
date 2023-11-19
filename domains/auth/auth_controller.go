package auth

import (
	"net/http"

	"github.com/0xfzz/tuwit-backend/domains/auth/dto"
	"github.com/0xfzz/tuwit-backend/lib"
	"github.com/0xfzz/tuwit-backend/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	api     lib.Api
	authSvc AuthService
}

func NewAuthController(api lib.Api, authSvc AuthService) AuthController {
	return AuthController{
		api,
		authSvc,
	}
}
func (c AuthController) login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response.Message(ctx, http.StatusBadRequest, "Please post data correctly wkwkwk")
		return
	}
	token, err := c.authSvc.LoginUser(loginRequest.Identifier, loginRequest.Password)
	if err != nil {
		response.Message(ctx, http.StatusBadRequest, "Username/Email or Password is wrong")
		return
	}
	responseToken := response.ResponseWithData{
		Data: gin.H{
			"access_token": token,
		},
		Message: "Login successfully",
	}
	response.Data(ctx, http.StatusOK, responseToken)
}
func (c AuthController) register(ctx *gin.Context) {
	var registerReq dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		response.Message(ctx, http.StatusBadRequest, "U forgot something")
		return
	}
	if registerReq.Password != registerReq.ConfirmPassword {
		response.Message(ctx, http.StatusBadRequest, "Confirm password must be same as password.")
		return
	}
	if c.authSvc.CheckUsernameAvailability(registerReq.Username) {
		response.Message(ctx, http.StatusBadRequest, "Username already exist.")
		return
	}
	var err = c.authSvc.RegisterUser(registerReq)
	if err != nil {
		response.InternalServerError(ctx)
	}
	response.Message(ctx, http.StatusOK, "User Created :).")
}

func (c AuthController) AddRoute() {
	authGroup := c.api.Route.Group("/auth")
	authGroup.POST("/login", c.login)
	authGroup.GET("/register", c.register)
}
