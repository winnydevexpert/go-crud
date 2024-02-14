package routers

import (
	"go-crud-learn/database"

	handler "go-crud-learn/feature/customer/delivery/http/v1"
	repo "go-crud-learn/feature/customer/repository"
	"go-crud-learn/feature/customer/usecase"

	"github.com/gin-gonic/gin"
)

func UserEndpoints(version *gin.RouterGroup) {
	router := version.Group("/customers")

	newRepo := repo.NewCustomerRepository(database.DB)
	newUseCase := usecase.NewCustomerUsecase(newRepo)

	newHandler := handler.NewCustomerHandler(newUseCase)

	router.POST("", newHandler.CreateCustomer)

	// router :=
	// version.POST("/customers", controllers.Register)
}
