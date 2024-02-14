package http

import (
	"go-crud-learn/constant"
	"go-crud-learn/domain"
	"go-crud-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	requestForm "go-crud-learn/domain/requestForm"
)

type customerHandler struct {
	customerUsecase domain.CustomerUsecase
}

func NewCustomerHandler(customerUsecase domain.CustomerUsecase) *customerHandler {
	return &customerHandler{
		customerUsecase,
	}

}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var req requestForm.CreateCustomerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), nil))
		return
	}

	customer, errCreateCustomer := h.customerUsecase.CreateCustomer(req)
	if errCreateCustomer != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorMessage(constant.ERROR_INTERNAL_SERVER_ERROR, nil))
	}

	title := "Customer"
	description := "Create Customer Successfully."
	c.JSON(http.StatusCreated, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        customer,
	}))

}
