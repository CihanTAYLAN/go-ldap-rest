package ldap_connector

import (
	"log"

	"github.com/go-ldap/ldap/v3"
)

type ConnectParams struct {
	LdapURL      string `json:"ldapURL" bson:"searchBase" binding:"required"`
	BindDN       string `json:"bindDN" bson:"searchFilter" binding:"required"`
	BindPassword string `json:"bindPassword" bson:"attributes" binding:"required"`
}

func Connect(ConnectParams ConnectParams) (*ldap.Conn, error) {
	l, err := ldap.DialURL(ConnectParams.LdapURL)
	if err != nil {
		return nil, err
	}
	err = l.Bind(ConnectParams.BindDN, ConnectParams.BindPassword)
	if err != nil {
		return nil, err
	}
	return l, nil
}

type FindParams struct {
	Conn         *ldap.Conn `json:"conn" bson:"conn" binding:"required"`
	SearchBase   string     `json:"searchBase" bson:"searchBase" binding:"required"`
	SearchFilter string     `json:"searchFilter" bson:"searchFilter" binding:"required"`
	Attributes   []string   `json:"attributes" bson:"attributes" default:"[\"*\"]"`
}

func Find(
	FindParams FindParams,
) []*ldap.Entry {
	if FindParams.Attributes == nil {
		FindParams.Attributes = []string{"*"}
	}

	searchRequest := ldap.NewSearchRequest(
		FindParams.SearchBase,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		FindParams.SearchFilter,
		FindParams.Attributes,
		nil,
	)
	sr, err := FindParams.Conn.SearchWithPaging(searchRequest, 4)
	if err != nil {
		log.Fatal(err)
	}
	return sr.Entries

}
