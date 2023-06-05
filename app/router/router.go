package router

import (
	"io/ioutil"
	"ldap-rest/app/controllers/ctrl_admin"
	"ldap-rest/docs"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	data, _ := ioutil.ReadFile("docs/swagger.json")
	ioutil.WriteFile("static/swagger.json", data, 0777)

	println(string(data))
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/doc/*any", func(ctx *gin.Context) { // Stoplight Documentation
		ctx.HTML(200, "stoplight.html", gin.H{})
	})

	docs.SwaggerInfo.Title = "Go Ldap Rest API"
	docs.SwaggerInfo.Host = "localhost:8088"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/users", ctrl_admin.Find)
		}
	}

	return r
}
