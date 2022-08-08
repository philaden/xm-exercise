package domains

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

type (
	User struct {
		ID          string `gorm:"primaryKey"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   *time.Time `gorm:"index"`
		FirstName   string
		LastName    string
		PhoneNumber string
		Email       string
		Password    string
	}
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}

func ToDtoUserResponse(user *User) *dto.UserResponseDto {
	return &dto.UserResponseDto{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
}

func ToSliceUserResponseDto(users []User) dto.UsersResponseDto {
	var dtos []dto.UserResponseDto
	for _, value := range users {
		dtos = append(dtos, dto.UserResponseDto{
			ID:          value.ID,
			FirstName:   value.FirstName,
			LastName:    value.LastName,
			PhoneNumber: value.PhoneNumber,
			Email:       value.Email,
		})
	}
	return dtos
}
