package main

import (
	"context"
	"fmt"
	config "go-crud-learn/config"
	"go-crud-learn/constant"
	"go-crud-learn/controllers"
	"go-crud-learn/database"
	"go-crud-learn/middlewares"
	"go-crud-learn/routers"

	"go-crud-learn/utils"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	_ "go-crud-learn/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectDB()
// }

func InitLoadEnv() {
	// Load config from env
	appConfig, errLoadConfig := config.LoadConfig("./config/env")
	if errLoadConfig != nil {
		log.Fatal("cannot load config:", errLoadConfig)
	}
	config.AppConfig = appConfig

	fmt.Println("----------------------------------")
	fmt.Println("MODE:", appConfig.Mode)
	fmt.Println("PORT:", appConfig.Port)
	fmt.Println("----------------------------------")
}

func InitDatabase() func() {
	closeDB, errConnectDB := database.ConnectDB(config.AppConfig.Mode)
	if errConnectDB != nil {
		utils.Logger.Fatal(errConnectDB.Error())
	}

	return closeDB
}

// @title base code
// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
// @version 1.0
func main() {

	app := gin.Default()

	app.SetTrustedProxies(nil)

	app.Use(middlewares.CORS())

	InitLoadEnv()

	InitDatabase()

	utils.InitializeLogger()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorMessage(constant.ERROR_ENDPOINT_NOT_FOUND, nil))
	})

	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, utils.ErrorMessage(constant.ERROR_METHOD_NOT_ALLOW, nil))
	})

	versionRouter := app.Group("/v1", middlewares.BasicMiddleware())

	routers.InitRouter(versionRouter)

	versionRouterV2 := app.Group("/v2", middlewares.BasicMiddleware())

	routers.InitRouter(versionRouterV2)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	errRunApp := app.Run(":8000")
	if errRunApp != nil {
		panic(errRunApp)

	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:         ":" + config.AppConfig.Port,
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if errListen := srv.ListenAndServe(); errListen != nil && errListen != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", errListen)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down ....")

	timeOutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if errShutdown := srv.Shutdown(timeOutctx); errShutdown != nil {
		utils.Logger.Error(errShutdown.Error())
	}

	// app.POST("/customers", controllers.Register)
	// app.POST("/orders", controllers.CreateOrder)
	// app.POST("/bills", controllers.CreateBill)
	app.GET("/bills/:id", controllers.GetBillById)
}
