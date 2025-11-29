package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/alirzamehrzad/drivey/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHnadler()

	r.GET("/", handler.Test)
	r.GET("/users", handler.Users)
	r.GET("/user/:id", handler.UserById)
	r.GET("/user/get-user-by-username/:username", handler.UserByUsername)
	r.GET("/user/:id/accounts", handler.Accounts)

	r.POST("/add-user", handler.AddUser)
	r.POST("/binder/header1", handler.HeaderBinder1)
	r.POST("/binder/header2", handler.HeaderBinder2)

	r.POST("/binder/query", handler.QueryBinder)

	r.POST("/binder/body", handler.BodyBinder)
	r.POST("/binder/form", handler.FormBinder)
	r.POST("/binder/file", handler.FileBinder)

}
