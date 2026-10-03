package response

import (
	"github.com/gin-gonic/gin"
)

// Response is a standardized API response envelope
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JSON sends a standardized JSON response
func JSON(c *gin.Context, statusCode int, data interface{}, err string) {
	success := statusCode >= 200 && statusCode < 300

	c.JSON(statusCode, Response{
		Success: success,
		Data:    data,
		Error:   err,
	})
}
