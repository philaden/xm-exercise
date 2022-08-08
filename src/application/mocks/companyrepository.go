// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/philaden/xm-go-challenge/src/application/repositories (interfaces: ICompanyRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domains "github.com/philaden/xm-go-challenge/src/application/domains"
)

// MockICompanyRepository is a mock of ICompanyRepository interface.
type MockICompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompanyRepositoryMockRecorder
}

// MockICompanyRepositoryMockRecorder is the mock recorder for MockICompanyRepository.
type MockICompanyRepositoryMockRecorder struct {
	mock *MockICompanyRepository
}

// NewMockICompanyRepository creates a new mock instance.
func NewMockICompanyRepository(ctrl *gomock.Controller) *MockICompanyRepository {
	mock := &MockICompanyRepository{ctrl: ctrl}
	mock.recorder = &MockICompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICompanyRepository) EXPECT() *MockICompanyRepositoryMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockICompanyRepository) CreateCompany(arg0, arg1, arg2, arg3, arg4 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockICompanyRepositoryMockRecorder) CreateCompany(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockICompanyRepository)(nil).CreateCompany), arg0, arg1, arg2, arg3, arg4)
}

// DeleteCompany mocks base method.
func (m *MockICompanyRepository) DeleteCompany(arg0 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockICompanyRepositoryMockRecorder) DeleteCompany(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockICompanyRepository)(nil).DeleteCompany), arg0)
}

// GetCompanies mocks base method.
func (m *MockICompanyRepository) GetCompanies(arg0 url.Values) ([]domains.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanies", arg0)
	ret0, _ := ret[0].([]domains.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanies indicates an expected call of GetCompanies.
func (mr *MockICompanyRepositoryMockRecorder) GetCompanies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanies", reflect.TypeOf((*MockICompanyRepository)(nil).GetCompanies), arg0)
}

// GetCompanyByCode mocks base method.
func (m *MockICompanyRepository) GetCompanyByCode(arg0 string) (*domains.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByCode", arg0)
	ret0, _ := ret[0].(*domains.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByCode indicates an expected call of GetCompanyByCode.
func (mr *MockICompanyRepositoryMockRecorder) GetCompanyByCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByCode", reflect.TypeOf((*MockICompanyRepository)(nil).GetCompanyByCode), arg0)
}