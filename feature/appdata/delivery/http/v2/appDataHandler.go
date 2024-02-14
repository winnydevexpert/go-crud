package http

import (
	constant "go-crud-learn/constant"
	"go-crud-learn/domain"
	"go-crud-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appDataHandler struct {
	appDataUsecase domain.AppDataUsecase
}

func NewAppDataHandler(appDataUsecase domain.AppDataUsecase) *appDataHandler {
	return &appDataHandler{
		appDataUsecase,
	}
}

// @Accept  json
// @Tags App data
// @Router /v2/appData/list [get]
// @Success 200 {object} domain.AppDataListResponse
func (h *appDataHandler) GetListAppData(c *gin.Context) {
	appDatas, errGetData := h.appDataUsecase.GetListAppData()
	if errGetData != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorMessage(constant.ERROR_INTERNAL_SERVER_ERROR, nil))
		return
	}

	pagination := utils.Pagination(1, 10, 23)

	title := "App Data"
	description := "app data list."
	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title:       &title,
		Description: &description,
		Items:       appDatas,
		Pagination:  &pagination,
	}))
}
