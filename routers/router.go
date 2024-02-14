package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(version *gin.RouterGroup) {
	AppDataRouter(version)
	UserRouter(version)
	UserEndpoints(version)
	OrderEndpoints(version)
	BillEndpoints(version)
}

func InitRouterV2(version *gin.RouterGroup) {
	AppDataRouterV2(version)
}
