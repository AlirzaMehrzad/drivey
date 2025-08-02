package routers

import (
	"github.com/alirzamehrzad/drivey/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.GET("/getById/:id", handler.GetHealthById)
}
