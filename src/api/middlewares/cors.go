package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/alirzamehrzad/drivey/config"
)

func Cors(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Cors.AllowOrigins)
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Max-Age", "21600")
		ctx.Set("content-type", "application/json")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
	}
}
