package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	Firstname    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	Lastname     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	Mobilenumber string `json:"mobile_number" binding:"required,min=11,max=11"`
}

func NewTestHnadler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (h *TestHandler) UserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.JSON(http.StatusOK, gin.H{
		"result":   "Test",
		"username": username,
	})
}

func (h *TestHandler) Accounts(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}
func (h *TestHandler) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
	})
}

func (h *TestHandler) HeaderBinder1(ctx *gin.Context) {
	// *** method1
	he := header{}
	ctx.BindHeader(&he)

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"header": he,
	})
}

func (h *TestHandler) HeaderBinder2(ctx *gin.Context) {
	// *** method2
	userId := ctx.GetHeader("userId")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"userId": userId,
	})
}

func (h *TestHandler) QueryBinder(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder",
		"id":     ids,
		"name":   name,
	})
}

func (h *TestHandler) BodyBinder(ctx *gin.Context) {
	p := personData{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *TestHandler) FormBinder(ctx *gin.Context) {
	p := personData{}
	ctx.ShouldBind(&p)

	ctx.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": p,
	})
}

func (h *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "fileDestination")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": file.Filename,
	})
}
