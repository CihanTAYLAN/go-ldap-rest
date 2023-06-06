package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}
