// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package network is a generated GoMock package.
package network

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockService) Register(root hash.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", root)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockServiceMockRecorder) Register(root interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockService)(nil).Register), root)
}
