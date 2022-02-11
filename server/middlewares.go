package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/go-jwt-sample-app/libraries"
)

func AuthorizeJWT() gin.HandlerFunc {

	jwtLib := libraries.NewJWTLib()

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "access token is required"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtLib.ValidateToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "invalid access token"})
			return
		}

		claims := token.Claims
		fmt.Println(claims)
	}
}
