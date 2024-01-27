// Code generated by MockGen. DO NOT EDIT.
// Source: internal/entity/temperature.go
//
// Generated by this command:
//
//	mockgen -source=internal/entity/temperature.go -destination=internal/usecase/temperature/mock/find_temperature.go
//

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	context "context"
	reflect "reflect"

	temperature "github.com/GeovaneCavalcante/temperatura-cep/pkg/temperature"
	gomock "go.uber.org/mock/gomock"
)

// MockFindTemperatureUseCase is a mock of FindTemperatureUseCase interface.
type MockFindTemperatureUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockFindTemperatureUseCaseMockRecorder
}

// MockFindTemperatureUseCaseMockRecorder is the mock recorder for MockFindTemperatureUseCase.
type MockFindTemperatureUseCaseMockRecorder struct {
	mock *MockFindTemperatureUseCase
}

// NewMockFindTemperatureUseCase creates a new mock instance.
func NewMockFindTemperatureUseCase(ctrl *gomock.Controller) *MockFindTemperatureUseCase {
	mock := &MockFindTemperatureUseCase{ctrl: ctrl}
	mock.recorder = &MockFindTemperatureUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindTemperatureUseCase) EXPECT() *MockFindTemperatureUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockFindTemperatureUseCase) Execute(ctx context.Context, zipCode string) (*temperature.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, zipCode)
	ret0, _ := ret[0].(*temperature.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockFindTemperatureUseCaseMockRecorder) Execute(ctx, zipCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockFindTemperatureUseCase)(nil).Execute), ctx, zipCode)
}
