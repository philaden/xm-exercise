package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	dto "github.com/philaden/xm-go-challenge/src/application/dtos"
	mocks "github.com/philaden/xm-go-challenge/src/application/mocks"
	"github.com/stretchr/testify/require"
)

func Test_GetCompanies(t *testing.T) {

	companies := mocks.GetMockCompanies()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().GetCompanies(nil).Return(companies, nil)

	var service ICompanyService = CompanyService{CompanyRepository: companyRepository}
	result, message, err := service.GetCompanies(nil)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Contains(t, message, "Operation Successful")
}

func Test_GetCompany_By_Valid_Code(t *testing.T) {

	const code string = "da8ef851e075"

	company := mocks.GetMockCompanyByStringParameter(code)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().GetCompanyByCode(code).Return(company, nil)

	var service ICompanyService = CompanyService{CompanyRepository: companyRepository}
	result, message, err := service.GetCompanyByCode(code)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Contains(t, message, "Operation Successful")
}

func Test_CreateCompany_With_Valid_Params(t *testing.T) {
	data := mocks.CreateMockCompanyPayload()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().CreateCompany(data.Name, data.Code, data.Country, data.Website, data.PhoneNumber).Return(true, nil)

	var service ICompanyService = CompanyService{CompanyRepository: companyRepository}

	registrationData := dto.CreateCompanyDto{
		Name:        data.Name,
		Code:        data.Code,
		Country:     data.Country,
		Website:     data.Website,
		PhoneNumber: data.PhoneNumber,
	}

	result, message, err := service.RegisterCompany(registrationData)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Contains(t, message, "registered successfully")
}

func Test_DeleteCompany_By_Id(t *testing.T) {

	const id uint = 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().DeleteCompany(int(id)).Return(true, nil)

	var service ICompanyService = CompanyService{CompanyRepository: companyRepository}

	_, _, err := service.DeleteCompany(int(id))

	require.NoError(t, err)
}
