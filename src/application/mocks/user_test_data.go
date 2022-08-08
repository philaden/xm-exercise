package mocks

import (
	"time"

	domain "github.com/philaden/xm-go-challenge/src/application/domains"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

func GetMockUsersResponseDto() dto.UsersResponseDto {
	mock := []dto.UserResponseDto{
		{
			ID:          "69ff4588-392a-4378-80ef-05fa21f53871",
			FirstName:   "Test",
			LastName:    "Dev",
			Email:       "test_dev@gmail.com",
			PhoneNumber: "99010101221",
		},
		{
			ID:          "df85e3b4-f04a-4e03-8f59-5bd5c550037b",
			FirstName:   "Tester",
			LastName:    "Dev",
			Email:       "tester_dev@gmail.com",
			PhoneNumber: "99010101211",
		},
		{
			ID:          "df85e3b4-f04a-4e03-8f59-5bd5c5500377",
			FirstName:   "Beta",
			LastName:    "Tester",
			Email:       "beta_tester@gmail.com",
			PhoneNumber: "99010101210",
		},
	}

	return mock
}

func GetMockUsersDto() []domain.User {
	mock := []domain.User{
		{
			ID:          "69ff4588-392a-4378-80ef-05fa21f53871",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			FirstName:   "Test",
			LastName:    "Dev",
			Email:       "test_dev@gmail.com",
			PhoneNumber: "99010101221",
		},
		{
			ID:          "df85e3b4-f04a-4e03-8f59-5bd5c550037b",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			FirstName:   "Tester",
			LastName:    "Dev",
			Email:       "tester_dev@gmail.com",
			PhoneNumber: "99010101211",
		},
		{
			ID:          "df85e3b4-f04a-4e03-8f59-5bd5c5500377",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			FirstName:   "Beta",
			LastName:    "Tester",
			Email:       "beta_tester@gmail.com",
			PhoneNumber: "99010101210",
		},
	}
	return mock
}

func CreateMockUserPayload() dto.CreateUserDto {
	return dto.CreateUserDto{
		FirstName:   "Beta",
		LastName:    "Tester",
		Email:       "beta_tester@gmail.com",
		PhoneNumber: "99010101210",
		Password:    "Password@1",
	}
}
