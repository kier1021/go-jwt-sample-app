package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-jwt-sample-app/api/dto"
	"github.com/kier1021/go-jwt-sample-app/api/services"
	"github.com/kier1021/go-jwt-sample-app/libraries"
)

type AuthController struct {
	jwtLib      *libraries.JWTLib
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		jwtLib:      libraries.NewJWTLib(),
		authService: authService,
	}
}

func (ctrl *AuthController) Login() gin.HandlerFunc {

	return func(c *gin.Context) {

		var loginCred dto.LoginRequestDTO

		err := c.ShouldBind(&loginCred)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": "error in data input"})
			return
		}

		if !ctrl.authService.Login(loginCred.Username, loginCred.Password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "username/password is incorrect"})
			return
		}

		token, err := ctrl.jwtLib.GenerateToken(map[string]interface{}{
			"user_id": loginCred.Username,
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"error": "error in generating access token"})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message":      "Login successfully",
			"access_token": token,
		})
	}
}
