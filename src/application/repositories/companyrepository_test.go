package repositories

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/philaden/xm-go-challenge/src/application/mocks"
	"github.com/stretchr/testify/require"
)

func Test_GetCompanies(t *testing.T) {

	companies := mocks.GetMockCompanies()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().GetCompanies(nil).Return(companies, nil)

	response, err := companyRepository.GetCompanies(nil)
	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, response, companies)
}

func Test_GetCompany_By_Valid_Code(t *testing.T) {

	const code string = "da8ef851e075"

	company := mocks.GetMockCompanyByStringParameter(code)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().GetCompanyByCode(code).Return(company, nil)

	response, err := companyRepository.GetCompanyByCode(code)
	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, response, company)
}

func Test_CreateCompany_With_Valid_Params(t *testing.T) {
	data := mocks.CreateMockCompanyPayload()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().CreateCompany(data.Name, data.Code, data.Country, data.Website, data.PhoneNumber).Return(true, nil)

	response, err := companyRepository.CreateCompany(data.Name, data.Code, data.Country, data.Website, data.PhoneNumber)
	require.NoError(t, err)
	require.Equal(t, response, true)
}

func Test_DeleteCompany_By_Id(t *testing.T) {

	const id uint = 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyRepository := mocks.NewMockICompanyRepository(ctrl)
	companyRepository.EXPECT().DeleteCompany(int(id)).Return(true, nil)

	response, err := companyRepository.DeleteCompany(int(id))
	require.NoError(t, err)
	require.Equal(t, response, true)
}
