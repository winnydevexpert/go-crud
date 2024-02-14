package usecase

import (
	"go-crud-learn/domain"
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
}

func NewCustomerUsecase(customerRepo domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (u customerUsecase) CreateCustomer(req requestForm.CreateCustomerRequest) (res responseForm.CreateCustomerResponse, err error) {
	var customerResponse requestForm.CreateCustomerRequest

	customerResponse, err = u.customerRepo.CreateCustomer(req)
	if err != nil {
		return res, err
	}
	res.FirstName = customerResponse.FirstName
	res.Body = customerResponse.Body
	return res, nil
}
