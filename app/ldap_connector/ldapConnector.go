package ldap_connector

import (
	"log"

	"github.com/go-ldap/ldap/v3"
)

// Host : ldap://ldap.forumsys.com
// User : cn=read-only-admin,dc=example,dc=com
// Pass : password

func Connect(ldapURL string, bindDN string, bindPassword string) *ldap.Conn {
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	err = l.Bind(bindDN, bindPassword)
	if err != nil {
		log.Fatal(err)
	}
	return l
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
