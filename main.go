package main

import (
	"Goland_CRUD/controllers"
	"Goland_CRUD/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
