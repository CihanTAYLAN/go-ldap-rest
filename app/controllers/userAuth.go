package controllers

import (
	"ldap-rest/app/helpers"
	"ldap-rest/app/ldap_connector"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthRequest struct {
	UserDn   string `json:"userDn" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserAuthResponse struct {
	helpers.ResponseType
	Data struct {
		auth bool
	}
}

// @Tags Ldap
// @Summary UserAuth
// @Description EG;<br>ByRG81IDDPQFY9+9dSaWFKIA3Xp1vZhrpCjCg4XXR7gnNxLM9SvgTK1PFKMrsdE5s4mNRSIo8qJhzeZAdMi5zQfAhJOV8FDdmEs=<br>username: cn=read-only-admin,dc=example,dc=com<br>password: password
// @ID UserAuth
// @Accept json
// @Produce json
// @Param Token header string true "Your Auth Token"
// @Param request body UserAuthRequest true "Request Body"
// @Success 200 {object} UserAuthResponse
// @Router /ldap/user-auth [post]
func UserAuth(c *gin.Context) {
	var request UserAuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	credentials, err := helpers.ReadSession(c.Request.Header["Token"][0])
	if err != nil {
		helpers.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
			"response": err,
		})
		return
	}

	ldapCon, err := ldap_connector.Connect(ldap_connector.ConnectParams{
		LdapURL:      credentials.LdapURL,
		BindDN:       credentials.BindDN,
		BindPassword: credentials.BindPassword,
	})

	badRequestMessage := "Bad Request"
	if err == nil {
		// First bind with a read only user
		err = ldapCon.Bind(request.UserDn, request.Password)
		if err != nil {
			helpers.CreateResponse(c, http.StatusBadRequest, badRequestMessage, gin.H{
				"response": err,
			})
			ldapCon.Close()
			return
		}

		// Bind as the user to verify their password
		err = ldapCon.Bind(request.UserDn, request.Password)
		if err != nil {
			helpers.CreateResponse(c, http.StatusBadRequest, badRequestMessage, gin.H{
				"response": err,
			})
			ldapCon.Close()
			return
		}

		// Rebind as the read only user for any further queries
		err = ldapCon.Bind(request.UserDn, request.Password)
		if err != nil {
			helpers.CreateResponse(c, http.StatusBadRequest, badRequestMessage, gin.H{
				"response": err,
			})
			ldapCon.Close()
			return
		}

		if err != nil {
			helpers.CreateResponse(c, http.StatusBadRequest, badRequestMessage, gin.H{
				"response": err,
			})
			ldapCon.Close()
			return
		}
		helpers.CreateResponse(c, http.StatusOK, "OK", gin.H{
			"auth": true,
		})
		ldapCon.Close()
		return
	}
	helpers.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
		"response": err,
	})
	return
}
