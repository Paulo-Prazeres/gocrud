package controllers

import (
	"github.com/Paulo-Prazeres/gocrud/initializers"
	"github.com/Paulo-Prazeres/gocrud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with the Posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id from the url

	id := c.Param("id")
	// Get the Posts
	var post models.Post
	initializers.DB.Find(&post, id)

	// Respond with the Posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id from the url

	id := c.Param("id")

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post to update
	var post models.Post
	initializers.DB.Find(&post, id)

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with the Posts
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id from the url

	id := c.Param("id")

	// Delete the post
	var post models.Post
	initializers.DB.Delete(&post, id)

	// Respond with the Posts
	c.Status(200)

}
