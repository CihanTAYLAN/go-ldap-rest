package helpers

import (
	"github.com/gin-gonic/gin"
)

func CreateResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, ResponseType{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

type ResponseType struct {
	Status  int
	Message string
	Data    interface{}
}
