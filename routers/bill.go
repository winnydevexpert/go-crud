package routers

import (
	"go-crud-learn/database"

	handler "go-crud-learn/feature/bill/delivery/http/v1"
	repo "go-crud-learn/feature/bill/repository"
	"go-crud-learn/feature/bill/usecase"

	"github.com/gin-gonic/gin"
)

func BillEndpoints(version *gin.RouterGroup) {
	router := version.Group("/bills")

	newRepo := repo.NewBillRepository(database.DB)
	newUseCase := usecase.NewBillUsecase(newRepo)

	newHandler := handler.NewBillHandler(newUseCase)

	router.POST("", newHandler.CreateBill)
}
