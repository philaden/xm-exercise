package services

import (
	"errors"

	domain "github.com/philaden/xm-go-challenge/src/application/domains"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
	"github.com/philaden/xm-go-challenge/src/application/helpers"
	repo "github.com/philaden/xm-go-challenge/src/application/repositories"
)

type (
	IUserService interface {
		CreateUser(model dto.CreateUserDto) (result string, message string, err error)
		GetUserByEmail(email string) (result *dto.UserResponseDto, message string, err error)
		GetUsers() (result dto.UsersResponseDto, message string, err error)
	}

	UserService struct {
		UserRepository repo.IUserRepository
	}
)

func NewUserService(userRepository repo.IUserRepository) UserService {
	return UserService{UserRepository: userRepository}
}

func (service UserService) CreateUser(model dto.CreateUserDto) (result string, message string, err error) {

	if !helpers.ValidateEmail(model.Email) {
		message = "Email provided is invalid, please try again."
		return "", message, errors.New(message)
	}

	if len(model.Password) <= 7 {
		message = "The specified password was to short, must be longer than 8 characters."
		return "", message, errors.New(message)
	}

	if !helpers.ContainsCapitalLetter(model.Password) {
		message = "The specified password does not contain a capital letter."
		return "", message, errors.New("The specified password does not contain a capital letter.")
	}

	if !helpers.ContainsSpecialCharacter(model.Password) {
		message = "The password must contain at least one special character."
		return "", message, errors.New(message)
	}

	result, err = service.UserRepository.CreateUser(model.FirstName, model.LastName, model.PhoneNumber, model.Email, model.Password)

	if err != nil {
		message = "Operation failed"
		return "", message, err
	}

	message = "Operation successful"
	return result, message, nil
}

func (service UserService) GetUserByEmail(email string) (result *dto.UserResponseDto, message string, err error) {
	var data *domain.User
	data, err = service.UserRepository.GetUserByEmail(email)

	if err != nil {
		message = "Operation failed"
		return nil, message, err
	}

	result = domain.ToDtoUserResponse(data)
	message = "Operation successful"
	return result, message, nil
}

func (service UserService) GetUsers() (result dto.UsersResponseDto, message string, err error) {
	var data []domain.User
	data, err = service.UserRepository.GetUsers()

	if err != nil {
		message = "Operation failed"
		return nil, message, err
	}
	result = domain.ToSliceUserResponseDto(data)
	message = "Operation successful"
	return result, message, nil
}
