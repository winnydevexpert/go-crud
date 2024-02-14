package routers

import (
	"go-crud-learn/database"

	handler "go-crud-learn/feature/appdata/delivery/http/v1"
	handlerV2 "go-crud-learn/feature/appdata/delivery/http/v2"
	repo "go-crud-learn/feature/appdata/repository"
	usecase "go-crud-learn/feature/appdata/usecase"

	"github.com/gin-gonic/gin"
)


func AppDataRouter(version *gin.RouterGroup) {
	router := version.Group("/appData")

	newRepo := repo.NewAppDataRepository(database.DB)
	newUsecase := usecase.NewAppDataUsecase(newRepo)

	newHandler := handler.NewAppDataHandler(newUsecase)

	router.GET("/list", newHandler.GetListAppData)

}

func AppDataRouterV2(version *gin.RouterGroup) {
	router := version.Group("/appData")

	newRepo := repo.NewAppDataRepository(database.DB)
	newUsecase := usecase.NewAppDataUsecase(newRepo)

	newHandler := handlerV2.NewAppDataHandler(newUsecase)

	router.GET("/list", newHandler.GetListAppData)

}
