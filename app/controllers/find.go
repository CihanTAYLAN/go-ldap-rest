package controllers

import (
	"ldap-rest/app/helpers"
	"ldap-rest/app/ldap_connector"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindRequest struct {
	SearchBase   string   `json:"search_base" bson:"search_base" binding:"required"`
	SearchFilter string   `json:"search_filter" bson:"search_filter" binding:"required"`
	Attributes   []string `json:"attributes" bson:"attributes"`
}

type FindResponse struct {
	helpers.ResponseType
	Data struct {
		SessionToken string `json:"session_token" bson:"session_token" binding:"required"`
	}
}

// @Tags Ldap
// @Summary Find
// @Description EG;<br>ByRG81IDDPQFY9+9dSaWFKIA3Xp1vZhrpCjCg4XXR7gnNxLM9SvgTK1PFKMrsdE5s4mNRSIo8qJhzeZAdMi5zQfAhJOV8FDdmEs=<br>SearchBase: dc=example,dc=com<br>SearchFilter: (objectClass=person)
// @ID Find
// @Accept json
// @Produce json
// @Param Token header string true "Your Auth Token"
// @Param request body FindRequest true "Request Body"
// @Success 200 {object} FindResponse
// @Router /ldap/find [post]
func Find(c *gin.Context) {
	var request FindRequest
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

	if err == nil {
		findRes, err := ldap_connector.Find(ldap_connector.FindParams{
			Conn:         ldapCon,
			SearchBase:   request.SearchBase,
			SearchFilter: request.SearchFilter,
			Attributes:   request.Attributes,
		})
		if err != nil {
			helpers.CreateResponse(c, http.StatusBadRequest, "Bad Request", gin.H{
				"response": err,
			})
			ldapCon.Close()
			return
		}
		helpers.CreateResponse(c, http.StatusOK, "OK", gin.H{
			"response": findRes,
		})
		ldapCon.Close()
		return
	}
	helpers.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{
		"response": err,
	})
	return
}
