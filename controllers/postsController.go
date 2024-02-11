package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//Get data from req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Body, Body: body.Title}
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//Get id of url
	id := c.Param(("id"))

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Response with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id of the url
	id := c.Param(("id"))

	//Get the data of req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//Response it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get the id of the url
	id := c.Param(("id"))

	//Delete this post by id
	initializers.DB.Delete(&models.Post{}, id)

	//Response
	c.Status(200)
}
