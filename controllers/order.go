package controllers

import (
	"go-crud-learn/initializers"
	"go-crud-learn/model"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	//Get data from req body
	var body struct {
		ProductName string
	}

	c.Bind(&body)

	//Create a post
	order := model.Order{
		ProductName: body.ProductName,
	}
	result := initializers.DB.Create(&order) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"order": order,
	})
}
