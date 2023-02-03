// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package http is a generated GoMock package.
package http

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClientDoer is a mock of ClientDoer interface.
type MockClientDoer struct {
	ctrl     *gomock.Controller
	recorder *MockClientDoerMockRecorder
}

// MockClientDoerMockRecorder is the mock recorder for MockClientDoer.
type MockClientDoerMockRecorder struct {
	mock *MockClientDoer
}

// NewMockClientDoer creates a new mock instance.
func NewMockClientDoer(ctrl *gomock.Controller) *MockClientDoer {
	mock := &MockClientDoer{ctrl: ctrl}
	mock.recorder = &MockClientDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientDoer) EXPECT() *MockClientDoerMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockClientDoer) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockClientDoerMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClientDoer)(nil).Do), arg0)
}
