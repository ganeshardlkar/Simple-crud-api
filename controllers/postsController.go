package controllers

import (
	"go_crud/initializers"
	"go_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	// get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create a new post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Get id off URL
	id := c.Param("id")

	// Fetch a post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
