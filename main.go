package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	app := gin.Default()
	app.POST("/posts", controllers.PostCreate)
	app.PUT("/posts/:id", controllers.PostUpdate)
	app.GET("/posts", controllers.GetAllPosts)
	app.GET("/posts/:id", controllers.PostShow)
	app.DELETE("/posts/:id", controllers.PostDelete)

	app.Run() // listen and serve on 0.0.0.0:3000
}