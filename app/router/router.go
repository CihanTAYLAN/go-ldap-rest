package router

import (
	"ldap-rest/app/controllers"
	"ldap-rest/docs"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/doc/*any", func(ctx *gin.Context) { // Stoplight Documentation
		ctx.HTML(200, "stoplight.html", gin.H{})
	})

	r.GET("/", controllers.Index)
	v1 := r.Group("/api/v1")
	{
		ldap := v1.Group("/ldap")
		{
			ldap.POST("login", controllers.Login)
			ldap.POST("find", controllers.Find)
			ldap.POST("user-auth", controllers.UserAuth)
		}
	}
	println(docs.SwaggerInfo.Title)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
