package domains

import (
	"time"

	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
)

type (
	Company struct {
		ID          uint `gorm:"primaryKey"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   *time.Time `gorm:"index"`
		Name        string
		Code        string
		Country     string
		Website     string
		PhoneNumber string
	}

	Companies []Company
)

func ToDto(company *Company) *dto.CompanyResponseDto {
	return &dto.CompanyResponseDto{
		ID:          company.ID,
		CreatedAt:   company.CreatedAt,
		Name:        company.Name,
		Code:        company.Code,
		Country:     company.Country,
		Website:     company.Website,
		PhoneNumber: company.PhoneNumber,
	}
}

func ToSliceDto(companies Companies) dto.CompaniesResponseDto {
	var dtos []dto.CompanyResponseDto
	for _, value := range companies {
		dtos = append(dtos, dto.CompanyResponseDto{
			ID:          value.ID,
			CreatedAt:   value.CreatedAt,
			Name:        value.Name,
			Code:        value.Code,
			Country:     value.Country,
			Website:     value.Website,
			PhoneNumber: value.PhoneNumber,
		})
	}
	return dtos
}
