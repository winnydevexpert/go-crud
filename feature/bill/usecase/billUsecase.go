package usecase

import (
	"fmt"
	"go-crud-learn/domain"
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
	"go-crud-learn/model"
)

type billUsecase struct {
	billRepo domain.BillRepository
}

func NewBillUsecase(billRepo domain.BillRepository) domain.BillUsecase {
	return &billUsecase{
		billRepo: billRepo,
	}
}

func (u billUsecase) CreateBill(req requestForm.CreateBillRequest) (res responseForm.CreateBillResponse, err error) {
	var billResponse model.Bill

	billResponse, err = u.billRepo.CreateBill(req)
	fmt.Println("Usecase: ", billResponse)
	if err != nil {
		return res, err
	}
	res.ID = billResponse.ID
	res.CustomerId = billResponse.CustomerId
	// res.Customer = billResponse.Customer
	res.OrderId = billResponse.OrderId
	// res.Order = billResponse.Order
	res.Amount = billResponse.Amount
	fmt.Println("res ", res)
	return res, nil
}
