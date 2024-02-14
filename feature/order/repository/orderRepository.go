package repository

import (
	"go-crud-learn/domain"
	requestForm "go-crud-learn/domain/requestForm"
	"go-crud-learn/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{DB: db}
}

func (q orderRepository) CreateOrder(req requestForm.CreateOrderRequest) (res requestForm.CreateOrderRequest, err error) {
	var order model.Order
	err = q.DB.Model(&order).Debug().Create(&req).Error
	if err != nil {
		return req, err
	}
	return req, nil
}
