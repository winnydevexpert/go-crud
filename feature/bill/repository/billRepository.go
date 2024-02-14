package repository

import (
	"fmt"
	"go-crud-learn/domain"
	requestForm "go-crud-learn/domain/requestForm"
	"go-crud-learn/model"

	"gorm.io/gorm"
)

type billRepository struct {
	DB *gorm.DB
}

func NewBillRepository(db *gorm.DB) domain.BillRepository {
	return &billRepository{DB: db}
}

func (q billRepository) CreateBill(req requestForm.CreateBillRequest) (res model.Bill, err error) {
	var bill model.Bill
	bill.CustomerId = req.CustomerId
	bill.OrderId = req.OrderId
	bill.Amount = req.Amount
	fmt.Println("Repo", bill)
	err = q.DB.Debug().Create(&bill).Error
	if err != nil {
		return bill, err
	}
	return bill, nil
}

// func (q billRepository) GetTransactionBillById(id int) (res model.Bill, err error) {
// 	id := c.Param(("id"))
// 	q.DB.First(&Bill)
// 	db.First(&user)
// }
