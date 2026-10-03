package handler

import (
	"krushi-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns the health status of the service
func HealthCheck(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"status": "UP"}, "")
}
