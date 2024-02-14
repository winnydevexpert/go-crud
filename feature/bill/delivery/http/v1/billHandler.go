package http

import (
	"go-crud-learn/constant"
	"go-crud-learn/domain"
	"go-crud-learn/utils"
	"net/http"

	requestForm "go-crud-learn/domain/requestForm"

	"github.com/gin-gonic/gin"
)

type billHanadler struct {
	billUsecase domain.BillUsecase
}

func NewBillHandler(billUsecase domain.BillUsecase) *billHanadler {
	return &billHanadler{
		billUsecase,
	}
}

func (h *billHanadler) CreateBill(c *gin.Context) {
	var req requestForm.CreateBillRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), nil))
		return
	}
	bill, errCreateBill := h.billUsecase.CreateBill(req)
	if errCreateBill != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorMessage(constant.ERROR_INTERNAL_SERVER_ERROR, nil))
		return
	}

	title := "Bill"
	description := "Create Bill Successfully"
	c.JSON(http.StatusCreated, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        bill,
	}))
}
