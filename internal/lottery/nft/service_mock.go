// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package nft is a generated GoMock package.
package nft

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageGenerator is a mock of ImageGenerator interface.
type MockImageGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockImageGeneratorMockRecorder
}

// MockImageGeneratorMockRecorder is the mock recorder for MockImageGenerator.
type MockImageGeneratorMockRecorder struct {
	mock *MockImageGenerator
}

// NewMockImageGenerator creates a new mock instance.
func NewMockImageGenerator(ctrl *gomock.Controller) *MockImageGenerator {
	mock := &MockImageGenerator{ctrl: ctrl}
	mock.recorder = &MockImageGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageGenerator) EXPECT() *MockImageGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockImageGenerator) Generate(hash string) (*LotteryValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", hash)
	ret0, _ := ret[0].(*LotteryValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockImageGeneratorMockRecorder) Generate(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockImageGenerator)(nil).Generate), hash)
}
