package usecase

import (
	"go-crud-learn/domain"
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUsecase(orderRepo domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}

func (u orderUsecase) CreateOrder(req requestForm.CreateOrderRequest) (res responseForm.CreateOrderResponse, err error) {
	var orderResponse requestForm.CreateOrderRequest

	orderResponse, err = u.orderRepo.CreateOrder(req)
	if err != nil {
		return res, err
	}
	res.ID = orderResponse.ID
	res.ProductName = orderResponse.ProductName
	return res, nil
}
