package ctrl_admin

import (
	"ldap-rest/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Users
// @Schemes
// @Tags Get Users
// @Accept json
// @Produce json
// @Success 200 {{ status:200,message:'OK',data:users:[] }}} Get Users
// @Failure 500 {object} FooBarResponse
// @Router /admin/users [get]
func Find(c *gin.Context) {

	helpers.CreateResponse(c, http.StatusOK, "OK", gin.H{
		"users": "sea",
	})
}
