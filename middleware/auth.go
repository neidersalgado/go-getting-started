package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		utiltoken := utils.JWTTokenService{}
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			ctx.Abort()
			return
		}

		email, err := utiltoken.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("email", email)
		ctx.Next()
	}
}
