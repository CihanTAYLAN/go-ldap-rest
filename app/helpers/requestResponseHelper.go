package helpers

import (
	"github.com/gin-gonic/gin"
)

func CreateResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
