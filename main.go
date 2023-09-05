package main

import (
	"JWT_AUTH/controllers"
	"JWT_AUTH/initializers"
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
