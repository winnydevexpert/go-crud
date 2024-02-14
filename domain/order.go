package domain

import (
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
)

type OrderUsecase interface {
	CreateOrder(req requestForm.CreateOrderRequest) (res responseForm.CreateOrderResponse, err error)
}

type OrderRepository interface {
	CreateOrder(req requestForm.CreateOrderRequest) (res requestForm.CreateOrderRequest, err error)
}
