package services

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	domain "github.com/philaden/xm-go-challenge/src/application/domains"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
	repo "github.com/philaden/xm-go-challenge/src/application/repositories"
)

type (
	ICompanyService interface {
		RegisterCompany(requestDto dto.CreateCompanyDto) (result bool, message string, err error)
		DeleteCompany(companyId int) (result bool, message string, err error)
		GetCompanies(queryParams url.Values) (result dto.CompaniesResponseDto, message string, err error)
		GetCompanyByCode(code string) (result *dto.CompanyResponseDto, message string, err error)
	}

	CompanyService struct {
		CompanyRepository repo.ICompanyRepository
	}
)

func NewCompanyService(companyRepository repo.ICompanyRepository) CompanyService {
	return CompanyService{
		CompanyRepository: companyRepository,
	}
}

func (service CompanyService) RegisterCompany(requestDto dto.CreateCompanyDto) (result bool, message string, err error) {

	result, err = service.CompanyRepository.CreateCompany(requestDto.Name, requestDto.Code, requestDto.Country, requestDto.Website, requestDto.PhoneNumber)

	message = "Something went wrong while trying to process that, please try again."
	if err != nil {
		return false, message, err
	}

	if !result {
		message = "Company already exists"
		return false, message, errors.New(message)
	}

	message = "company registered successfully"
	return result, message, nil
}

func (service CompanyService) DeleteCompany(companyId int) (result bool, message string, err error) {

	response, err := service.CompanyRepository.DeleteCompany(companyId)

	if err == nil && response {
		message = fmt.Sprintf("Company with Id %d does not exist", companyId)
		return false, message, err
	} else if err != nil {
		message = "Unable to delete company"
		return response, message, err
	}

	message = "Operation Successful"
	return true, message, nil
}

func (service CompanyService) GetCompanies(queryParams url.Values) (result dto.CompaniesResponseDto, message string, err error) {
	data, err := service.CompanyRepository.GetCompanies(queryParams)
	if err != nil {
		message = "Something went wrong while trying to process that, please try again."
		return nil, message, err
	}

	message = "Operation Successful"
	return domain.ToSliceDto(data), message, nil
}

func (service CompanyService) GetCompanyByCode(code string) (result *dto.CompanyResponseDto, message string, err error) {
	data, err := service.CompanyRepository.GetCompanyByCode(code)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		message = fmt.Sprintf("Company with code %s does not exist", code)
		return nil, message, err
	} else if err != nil {
		message = "Something went wrong while trying to process that, please try again."
		return nil, message, err
	}

	message = "Operation Successful"
	return domain.ToDto(data), message, nil
}
