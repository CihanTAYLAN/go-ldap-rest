package controllers

import (
	"ldap-rest/app/helpers"
	"ldap-rest/app/ldap_connector"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Ldap
// @Summary Login
// @Description EG;<br>LdapURL: ldap://ldap.forumsys.com:389<br>BindDN: cn=read-only-admin,dc=example,dc=com<br>BindPassword: password
// @ID Login
// @Accept json
// @Produce json
// @Param request body ldap_connector.ConnectParams true "Request Body"
// @Success 200 {object} ldap_connector.ConnectParams
// @Router /ldap/login [post]
func Login(c *gin.Context) {
	var request ldap_connector.ConnectParams
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ldapCon, err := ldap_connector.Connect(ldap_connector.ConnectParams{
		LdapURL:      request.LdapURL,
		BindDN:       request.BindDN,
		BindPassword: request.BindPassword,
	})
	ldapCon.Close()
	if err == nil {
		hash, err := helpers.CreateSession(&helpers.CreateSessionParams{
			LdapURL:      request.LdapURL,
			BindDN:       request.BindDN,
			BindPassword: request.BindPassword,
		})
		if err == nil {
			helpers.CreateResponse(c, http.StatusOK, "OK", gin.H{
				"session_token": hash,
			})
			return
		}
		helpers.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
			"response": err,
		})
		return
	}
	helpers.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
		"response": err,
	})
	return
}