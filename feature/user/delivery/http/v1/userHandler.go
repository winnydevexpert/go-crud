package http

import (
	"go-crud-learn/constant"
	requestForm "go-crud-learn/domain/requestForm"
	responseForm "go-crud-learn/domain/responseForm"
	"go-crud-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

// @Accept  json
// @Tags User
// @Router /v1/user/register [post]
// @param Body	body domain.UserRegisterRequest true "User register"
// @Success 201  {object}  domain.UserRegisterResponse
// @Failure 400 {object} utils.ErrorMessagePrototype
func (h *userHandler) Register(c *gin.Context) {

	var body requestForm.UserRegisterRequest

	errBindBody := c.ShouldBind(&body)
	if errBindBody != nil {
		errMessage := errBindBody.Error()
		utils.Logger.Error(errMessage)
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.ERROR_BAD_REQUEST, &errMessage))
		return
	}
	var userResponse responseForm.UserRegister
	title := "User"
	description := "Create user is successfully"
	c.JSON(http.StatusCreated, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Item:        userResponse,
	}))
}
