package routers

import (
	"go-crud-learn/database"

	handler "go-crud-learn/feature/order/delivery/http/v1"
	repo "go-crud-learn/feature/order/repository"
	"go-crud-learn/feature/order/usecase"

	"github.com/gin-gonic/gin"
)

func OrderEndpoints(version *gin.RouterGroup) {
	router := version.Group("/orders")

	newRepo := repo.NewOrderRepository(database.DB)
	newUseCase := usecase.NewOrderUsecase(newRepo)

	newHandler := handler.NewOrderHandler(newUseCase)

	router.POST("", newHandler.CreateOrder)
}
