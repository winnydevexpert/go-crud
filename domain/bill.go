package domain

import (
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
	"go-crud-learn/model"
)

type BillUsecase interface {
	CreateBill(req requestForm.CreateBillRequest) (res responseForm.CreateBillResponse, err error)
}

type BillRepository interface {
	CreateBill(req requestForm.CreateBillRequest) (res model.Bill, err error)
}

// type TransactionBillUsecase interface {
// 	GetTransactionBillById(req requestForm.CreateBillRequest) (res responseForm.CreateBillResponse, err error)
// }

// type TransactionBillRepository interface {
// 	GetTransactionBillById(req requestForm.CreateCustomerRequest) (res requestForm.CreateCustomerRequest, err error)
// }
