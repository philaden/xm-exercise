package mocks

import (
	"time"

	domain "github.com/philaden/xm-go-challenge/src/application/domains"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

func GetMockCompaniesResponseDto() dto.CompaniesResponseDto {
	mock := []dto.CompanyResponseDto{
		{
			ID:          1,
			CreatedAt:   time.Now(),
			Name:        "ABC Limited",
			Code:        "da8ef851e075",
			Country:     "Cyprus",
			Website:     "https://abclimited.com",
			PhoneNumber: "99010101221",
		},
		{
			ID:          2,
			CreatedAt:   time.Now(),
			Name:        "EFG Associates",
			Code:        "da8ef851e076",
			Country:     "Cyprus",
			Website:     "https://efgassociates.com",
			PhoneNumber: "99010101220",
		},
		{
			ID:          3,
			CreatedAt:   time.Now(),
			Name:        "XYZ Global Enterprise",
			Code:        "da8ef851e077",
			Country:     "Nigeria",
			Website:     "https://efgassociates.com",
			PhoneNumber: "2348062523084",
		},
	}

	return mock
}

func GetMockCompanies() []domain.Company {
	mock := []domain.Company{
		{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Name:        "ABC Limited",
			Code:        "da8ef851e075",
			Country:     "Cyprus",
			Website:     "https://abclimited.com",
			PhoneNumber: "99010101221",
		},
		{
			ID:          2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Name:        "EFG Associates",
			Code:        "da8ef851e076",
			Country:     "Cyprus",
			Website:     "https://efgassociates.com",
			PhoneNumber: "99010101220",
		},
		{
			ID:          3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Name:        "XYZ Global Enterprise",
			Code:        "da8ef851e077",
			Country:     "Nigeria",
			Website:     "https://efgassociates.com",
			PhoneNumber: "2348062523084",
		},
	}
	return mock
}

func GetMockCompanyByStringParameter(code string) *domain.Company {
	if code == "" {
		return nil
	}
	companies := GetMockCompanies()
	for _, com := range companies {

		switch code {
		case com.Name, com.Country, com.Code, com.Website, com.PhoneNumber:
			return &com
		default:
			return nil
		}
	}
	return nil
}

func GetMockCompanyById(id uint) *domain.Company {
	if id == 0 {
		return nil
	}
	companies := GetMockCompanies()
	for _, com := range companies {
		if com.ID == id {
			return &com
		}
	}
	return nil
}

func CreateMockCompanyPayload() dto.CreateCompanyDto {
	return dto.CreateCompanyDto{
		Name:        "ABC Limited",
		Code:        "da8ef851e075",
		Country:     "Cyprus",
		Website:     "https://abclimited.com",
		PhoneNumber: "99010101221",
	}
}
