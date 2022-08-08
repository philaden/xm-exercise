package dtos

import (
	"time"
)

type (
	CreateCompanyDto struct {
		Name        string `json:"name" binding:"required"`
		Country     string `json:"country" binding:"required"`
		Code        string `json:"code" binding:"required"`
		Website     string `json:"website" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
	}

	CompanyResponseDto struct {
		ID          uint
		CreatedAt   time.Time
		Name        string
		Country     string
		Code        string
		Website     string
		PhoneNumber string
	}

	CompaniesResponseDto []CompanyResponseDto

	GetCompaniesRequestDto struct {
		Name        string `json:"name"`
		Country     string `json:"country"`
		Code        string `json:"code"`
		Website     string `json:"website"`
		PhoneNumber string `json:"phoneNumber"`
	}
)
