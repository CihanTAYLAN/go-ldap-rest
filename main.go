package main

import (
	"ldap-rest/app/router"
)

func main() {
	r := router.SetupRouter()

	r.Run()
	// r.Run(":8088")
}
