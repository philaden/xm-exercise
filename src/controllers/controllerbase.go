package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philaden/xm-go-challenge/src/application/helpers"
	middleWare "github.com/philaden/xm-go-challenge/src/application/middlewares"
	"github.com/philaden/xm-go-challenge/src/application/repositories"
	"github.com/philaden/xm-go-challenge/src/application/services"
	"github.com/philaden/xm-go-challenge/src/infrastructure"
)

type ApiService struct {
	CompanyService services.ICompanyService
	UserService    services.IUserService
	AuthService    services.IAuthService
}

func NewApiService() ApiService {

	dbConfig, err := infrastructure.LoadConfiguration(".")
	if err != nil {
		fmt.Println(err)
		panic("failed to load application configuration settings")
	}

	userRepo := repositories.NewUserRepository(infrastructure.Connection)

	companyRepo := repositories.NewCompanyRepostiory(infrastructure.Connection)

	return ApiService{
		CompanyService: services.NewCompanyService(companyRepo),
		UserService:    services.NewUserService(userRepo),
		AuthService:    services.NewAuthService(userRepo, dbConfig),
	}
}

func SetupContollerRoutes(router *gin.Engine, config infrastructure.AppConfiguration) {

	services := NewApiService()

	controller := router.Group("/api")

	controller.POST("/auth", services.HandleUserLogin)

	controller.POST("/companies", helpers.Authenticate(config.SecretKey), middleWare.LocationAuth(), services.HandleCompanyRegistration)
	controller.DELETE("/companies/:id", helpers.Authenticate(config.SecretKey), middleWare.LocationAuth(), services.HandleDeleteCompany)
	controller.GET("/companies", services.HandleGetCompanies)
	controller.GET("/companies/:code", services.HandleGetCompanyByCompanyCode)

	controller.POST("/users", services.HandleCreateUser)
	controller.GET("/users", helpers.Authenticate(config.SecretKey), services.GetUsers)
	controller.GET("/users/userbyemail", helpers.Authenticate(config.SecretKey), services.HandleGetUserByEmail)
}
