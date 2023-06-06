package ldap_connector

import (
	"github.com/go-ldap/ldap/v3"
)

type ConnectParams struct {
	LdapURL      string
	BindDN       string
	BindPassword string
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
	Conn         *ldap.Conn
	SearchBase   string
	SearchFilter string
	Attributes   []string
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

type LdapEntry struct {
	DN         string
	Attributes []*LdapEntryAttribute
}

type LdapEntryAttribute struct {
	// Name is the name of the attribute
	Name string
	// Values contain the string values of the attribute
	Values []string
	// ByteValues contain the raw values of the attribute
	ByteValues [][]byte
}
