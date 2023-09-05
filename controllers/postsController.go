package controllers

import (
	"JWT_AUTH/initializers"
	"JWT_AUTH/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "1st post title", Body: "1st post body"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
