// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/Alkanoid/domain (interfaces: TransactionValidationService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mniak/Alkanoid/domain"
)

// MockTransactionValidationService is a mock of TransactionValidationService interface.
type MockTransactionValidationService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionValidationServiceMockRecorder
}

// MockTransactionValidationServiceMockRecorder is the mock recorder for MockTransactionValidationService.
type MockTransactionValidationServiceMockRecorder struct {
	mock *MockTransactionValidationService
}

// NewMockTransactionValidationService creates a new mock instance.
func NewMockTransactionValidationService(ctrl *gomock.Controller) *MockTransactionValidationService {
	mock := &MockTransactionValidationService{ctrl: ctrl}
	mock.recorder = &MockTransactionValidationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionValidationService) EXPECT() *MockTransactionValidationServiceMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockTransactionValidationService) Validate(arg0 domain.Transaction) domain.ValidationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(domain.ValidationResult)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockTransactionValidationServiceMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTransactionValidationService)(nil).Validate), arg0)
}
