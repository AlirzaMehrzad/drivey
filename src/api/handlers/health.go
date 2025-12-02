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

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "working")
}

func (h *HealthHandler) HealthPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "working post")
}
