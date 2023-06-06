package main

import (
	"ldap-rest/app"
)

// @Title	Go Ldap Rest API
// @version	1.0
// @description	This is a go ldap rest API Documentation.
// @servers	http://localhost:8088/api/v1
// @contact.name   API Support
// @contact.url    https://cihantaylan.com
// @contact.email  cihantaylan@cihantaylan.com
// @host      localhost:8088
// @BasePath  /api/v1
func main() {
	app.Bootstrap()
}
