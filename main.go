package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	middleWare "github.com/philaden/xm-go-challenge/src/application/middlewares"
	controller "github.com/philaden/xm-go-challenge/src/controllers"
	"github.com/philaden/xm-go-challenge/src/infrastructure"
)

func main() {

	config, err := infrastructure.LoadConfiguration(".")

	if err != nil {
		fmt.Println(err)
		panic("failed to load application configuration settings")
	}

	port := fmt.Sprintf(":%d", config.AppPort)

	if port == ":" {
		port = ":8000"
	}

	infrastructure.SetUpDatabaseServices(config)

	router := gin.Default()

	router.Use(middleWare.Cors())

	RegisterController(router, config)

	if err := router.Run(port); err != nil {
		fmt.Print(err)
	}
}

func RegisterController(router *gin.Engine, config infrastructure.AppConfiguration) {
	controller.SetupContollerRoutes(router, config)
}
