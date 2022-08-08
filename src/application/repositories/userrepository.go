package repositories

import (
	"errors"

	"github.com/jinzhu/gorm"
	domain "github.com/philaden/xm-go-challenge/src/application/domains"
	"github.com/philaden/xm-go-challenge/src/application/helpers"
)

type (
	IUserRepository interface {
		CreateUser(firstName, lastName, phoneNumber, email, password string) (string, error)
		GetUserByEmail(email string) (*domain.User, error)
		GetUsers() ([]domain.User, error)
	}

	UserRepository struct {
		DbContext *gorm.DB
	}
)

func NewUserRepository(dbContext *gorm.DB) IUserRepository {
	return UserRepository{DbContext: dbContext}
}

func (repo UserRepository) CreateUser(firstName, lastName, phoneNumber, email, password string) (string, error) {

	var foundUsers []domain.User

	if err := repo.DbContext.Select("email").Where("email = ?", email).Find(&foundUsers).Error; err != nil {
		return "", err
	}

	if len(foundUsers) > 0 {
		return "", errors.New("A user with that email address already exists")
	}

	hash, hashErr := helpers.HashPassword([]byte(password))

	if hashErr != nil {
		return "", hashErr
	}

	var user = domain.User{FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Email: email, Password: hash}

	if err := repo.DbContext.Create(&user).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}

func (repo UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.DbContext.Where("email = ? ", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	if err := repo.DbContext.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
