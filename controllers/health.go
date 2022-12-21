package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}
