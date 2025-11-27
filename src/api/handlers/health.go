package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ***class
type HealthHandler struct {
}

// *** constructor
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// => handlers
func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "working")
}

func (h *HealthHandler) HealthPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "working post")
}
