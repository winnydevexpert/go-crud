package domain

import (
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
)

type CustomerUsecase interface {
	CreateCustomer(req requestForm.CreateCustomerRequest) (res responseForm.CreateCustomerResponse, err error)
}

type CustomerRepository interface {
	CreateCustomer(req requestForm.CreateCustomerRequest) (res requestForm.CreateCustomerRequest, err error)
}
