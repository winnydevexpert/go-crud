package routers

import (
	handler "go-crud-learn/feature/user/delivery/http/v1"

	"github.com/gin-gonic/gin"
)

func UserRouter(version *gin.RouterGroup) {
	router := version.Group("/user")

	newHandler := handler.NewUserHandler()

	router.POST("/register", newHandler.Register)

}
