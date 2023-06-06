package app

import (
	"io/ioutil"
	"ldap-rest/app/router"
)

func Bootstrap() {
	println("------------ Bootstrap Gin App ------------")

	data, _ := ioutil.ReadFile("docs/swagger.json")
	ioutil.WriteFile("static/swagger.json", data, 0777)
	r := router.SetupRouter()

	// LDAP Development

	// ldapCon := ldap_connector.Connect(ldap_connector.ConnectParams{
	// 	LdapURL:      "ldap://ldap.forumsys.com:389",
	// 	BindDN:       "cn=read-only-admin,dc=example,dc=com",
	// 	BindPassword: "password",
	// })
	// searchBase := "dc=example,dc=com"
	// searchFilter := "(objectClass=person)"
	// rows := ldap_connector.Find(ldap_connector.FindParams{
	// 	Conn:         ldapCon,
	// 	SearchBase:   searchBase,
	// 	SearchFilter: searchFilter,
	// })
	// ldapCon.Close()

	// for index, user := range rows {
	// 	log.Println(index, " ", user.GetAttributeValue("cn"))
	// }

	r.Run()
	// r.Run(":8088")
}
