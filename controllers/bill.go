package controllers

import (
	"go-crud-learn/initializers"
	"go-crud-learn/model"

	"github.com/gin-gonic/gin"
)

func CreateBill(c *gin.Context) {

	var body struct {
		CustomerId uint
		OrderId    uint
		Amount     int
	}

	c.Bind(&body)

	//Create a bill
	bill := model.Bill{
		CustomerId: body.CustomerId,
		OrderId:    body.OrderId,
		Amount:     body.Amount,
	}
	result := initializers.DB.Create(&bill)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"bill": bill,
	})

}

func GetBillById(c *gin.Context) {

	id := c.Param(("id"))
	bill := model.Bill{}

	initializers.DB.Model(&model.Bill{}).Debug().Select("bills.amount").
		InnerJoins("Customer", initializers.DB.Select("first_name")).
		InnerJoins("Order", initializers.DB.Select("id", "product_name")).Where("bills.customer_id = ?", id).
		Find(&bill)
	c.JSON(200, gin.H{
		"bill": bill,
	})
}
