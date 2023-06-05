package controllers

import (
	"ldap-rest/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	helpers.CreateResponse(c, http.StatusOK, "OK", gin.H{
		"s": "s",
	})
}
