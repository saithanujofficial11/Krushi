package handler

import (
	"krushi-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping returns a simple pong message
func Ping(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "pong"}, "")
}
