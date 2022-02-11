package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-jwt-sample-app/api/controllers"
	"github.com/kier1021/go-jwt-sample-app/api/services"
)

func NewRouter() (r *gin.Engine) {
	r = gin.New()

	setRoutes(r)

	return r
}

func setRoutes(r *gin.Engine) {

	authService := services.NewAuthService()
	authCtrl := controllers.NewAuthController(authService)

	r.POST(
		"/auth/login",
		authCtrl.Login(),
	)

	r.GET(
		"/",
		AuthorizeJWT(),
		func() gin.HandlerFunc {
			return func(c *gin.Context) {
				c.JSON(200, map[string]interface{}{"message": "Hello World!!"})
			}
		}(),
	)
}
