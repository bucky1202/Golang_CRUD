package main

import (
	"Goland_CRUD/initializers"
	"Goland_CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
