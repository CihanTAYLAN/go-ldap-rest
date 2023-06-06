package ldap_connector

import (
	"github.com/go-ldap/ldap/v3"
)

type ConnectParams struct {
	LdapURL      string `json:"ldap_url" bson:"ldap_url" binding:"required"`
	BindDN       string `json:"bind_dn" bson:"bind_dn" binding:"required"`
	BindPassword string `json:"bind_password" bson:"bind_password" binding:"required"`
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
	SearchBase   string     `json:"search_base" bson:"search_base" binding:"required"`
	SearchFilter string     `json:"search_filter" bson:"search_filter" binding:"required"`
	Attributes   []string   `json:"attributes" bson:"attributes" default:"[\"*\"]"`
}

func Find(
	FindParams FindParams,
) ([]*ldap.Entry, error) {
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
		return nil, err
	}
	return sr.Entries, nil

}
