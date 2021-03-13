package main

import (
	"./router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		c.Next()
	}, cors.Default())

	engine.GET("/l7", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "ok"})
	})

	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			//users.GET("")
			//users.PATCH("")
			users.POST("", router.CreateUser)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("", router.Login)
		}
	}

	engine.Run()
}
