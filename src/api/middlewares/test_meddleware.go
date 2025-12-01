package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("x-api-key")
		if apiKey == "1" {
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"result": "Api key is required",
		})
	}
}
