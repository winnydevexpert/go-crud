package main

import (
	"go-crud-learn/initializers"
	"go-crud-learn/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	 // Migrate the schema
	 initializers.DB.AutoMigrate(&model.Customer{})
	 initializers.DB.AutoMigrate(&model.Order{})
	 initializers.DB.AutoMigrate(&model.Bill{})
}