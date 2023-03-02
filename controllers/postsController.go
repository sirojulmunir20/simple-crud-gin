package controllers

import (
	"go-crud-robby/initializers"
	"go-crud-robby/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// GET data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with then
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with then
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")

	// Get the data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"message": "Delete success",
	})
}
