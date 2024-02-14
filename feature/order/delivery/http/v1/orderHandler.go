package http

import (
	"go-crud-learn/constant"
	"go-crud-learn/domain"
	"go-crud-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	requestForm "go-crud-learn/domain/requestForm"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(orderUsecase domain.OrderUsecase) *orderHandler {
	return &orderHandler{
		orderUsecase,
	}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var req requestForm.CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), nil))
		return
	}

	order, errCreateOrder := h.orderUsecase.CreateOrder(req)
	if errCreateOrder != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorMessage(constant.ERROR_INTERNAL_SERVER_ERROR, nil))
		return
	}

	title := "Order"
	description := "Create Order Successfully"
	c.JSON(http.StatusCreated, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        order,
	}))
}
