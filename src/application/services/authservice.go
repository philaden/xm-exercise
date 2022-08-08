package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
	"github.com/philaden/xm-go-challenge/src/application/helpers"
	repo "github.com/philaden/xm-go-challenge/src/application/repositories"
	"github.com/philaden/xm-go-challenge/src/infrastructure"
)

type (
	IAuthService interface {
		Login(model dto.LoginDto) (result *dto.TokenDto, message string, err error)
	}

	AuthService struct {
		UserRepository repo.IUserRepository
		AppConfig      infrastructure.AppConfiguration
	}
)

func NewAuthService(userRepository repo.IUserRepository, appConfig infrastructure.AppConfiguration) AuthService {
	return AuthService{UserRepository: userRepository, AppConfig: appConfig}
}

func (service AuthService) Login(model dto.LoginDto) (result *dto.TokenDto, message string, err error) {

	if !helpers.ValidateEmail(model.Email) {
		message = "Email provided is invalid, please try again."
		return nil, message, errors.New(message)
	}

	user, err := service.UserRepository.GetUserByEmail(model.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		message = "account does not exist. Kindly sign up"
		return nil, message, err
	} else if err != nil {
		message := "Unable to login"
		return nil, message, err
	}

	if matched := helpers.CheckPasswordHash(model.Password, user.Password); matched != true {
		message = "passwords did not match"
		return nil, message, errors.New(message)
	}

	response, err := helpers.CreateSimpleToken(model.Email, user.ID, service.AppConfig.SecretKey)
	if err != nil {
		message = "Unable to generate token"
		return nil, message, err
	}

	message = "login successful"
	result = &dto.TokenDto{
		Email:  user.Email,
		Token:  response,
		UserId: user.ID,
	}

	return result, message, nil
}
