package controllers

import (
	"net/http"

	"github.com/andsholinka/golang-gin/initializers"
	"github.com/andsholinka/golang-gin/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to API TODO",
	})
}

func PostsCreate(c *gin.Context) {
	// Get data off re body
	var body struct {
		Title string
		Email string
	}

	c.Bind(&body)

	// Validate request body
	if body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Email: body.Email}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Limit(3).Find(&posts)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get data off re body
	var body struct {
		Title string
		Email string
	}

	c.Bind(&body)

	// Find the data were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Validate

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Email: body.Email,
	})

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Delete the data
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
