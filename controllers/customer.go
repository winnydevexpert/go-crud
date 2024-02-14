package controllers

import (
	"go-crud-learn/initializers"
	"go-crud-learn/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	//Get data from req body
	var body struct {
		FirstName string
		Body string
	}

	c.Bind(&body)

	//Create a post
	user := model.Customer{
		FirstName: body.FirstName,
		Body: body.Body,
	}
	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}