package repository

import (
	"go-crud-learn/domain"
	"go-crud-learn/model"

	requestForm "go-crud-learn/domain/requestForm"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{DB: db}
}

func (q customerRepository) CreateCustomer(req requestForm.CreateCustomerRequest) (res requestForm.CreateCustomerRequest, err error) {
	var customer model.Customer
	err = q.DB.Model(&customer).Create(&req).Error

	if err != nil {
		return req, err
	}
	return req, nil
}
