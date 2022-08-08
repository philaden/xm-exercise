package repositories

import (
	"net/url"

	"github.com/jinzhu/gorm"
	domain "github.com/philaden/xm-go-challenge/src/application/domains"
)

type (
	ICompanyRepository interface {
		CreateCompany(name, code, country, website, phoneNumber string) (bool, error)
		DeleteCompany(id int) (bool, error)
		GetCompanies(filters url.Values) ([]domain.Company, error)
		GetCompanyByCode(code string) (*domain.Company, error)
	}

	CompanyRepository struct {
		DbContext *gorm.DB
	}
)

func NewCompanyRepostiory(dbContext *gorm.DB) ICompanyRepository {
	return CompanyRepository{DbContext: dbContext}
}

func (repo CompanyRepository) CreateCompany(name, code, country, website, phoneNumber string) (bool, error) {

	var com *domain.Company = &domain.Company{}

	err := repo.DbContext.Where(&domain.Company{Code: code}).First(&com).Error

	if err == gorm.ErrRecordNotFound {
		newCompany := domain.Company{
			Name:        name,
			Code:        code,
			Country:     country,
			Website:     website,
			PhoneNumber: phoneNumber,
		}

		if err := repo.DbContext.Create(&newCompany).Error; err == nil {
			return true, nil
		}
	}
	return false, err
}

func (repo CompanyRepository) DeleteCompany(id int) (bool, error) {

	err := repo.DbContext.Delete(&domain.Company{}, id).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func (repo CompanyRepository) GetCompanies(filters url.Values) ([]domain.Company, error) {

	var companies []domain.Company

	params := make(map[string]interface{})

	for key, value := range filters {
		params[key] = value[0]
	}

	if err := repo.DbContext.Where(params).Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (repo CompanyRepository) GetCompanyByCode(code string) (*domain.Company, error) {

	com := domain.Company{}
	if err := repo.DbContext.Where(&domain.Company{Code: code}).First(&com).Error; err != nil {
		return nil, err
	}
	return &com, nil
}
