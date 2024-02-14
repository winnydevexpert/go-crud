package domain

import "go-crud-learn/model"

type CreateBillRequest struct {
	ID         uint `json:"id"`
	CustomerId uint `json:"customerId"`
	// Customer   model.Customer `gorm:"foreignKey:CustomerId" json:"customer"`
	OrderId uint `json:"orderId"`
	// Order      model.Order    `gorm:"foreignKey:OrderId" json:"order"`
	Amount int `json:"amount"`
}

type GetTransactionBillRequest struct {
	ID         uint           `json:"id"`
	CustomerId uint           `json:"customerId"`
	Customer   model.Customer `gorm:"foreignKey:CustomerId" json:"customer"`
	OrderId    uint           `json:"orderId"`
	Order      model.Order    `gorm:"foreignKey:OrderId" json:"order"`
	Amount     int            `json:"amount"`
}
