// Code generated by MockGen. DO NOT EDIT.
// Source: register.go

// Package lottery is a generated GoMock package.
package lottery

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
)

// MockRegister is a mock of Register interface.
type MockRegister struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterMockRecorder
}

// MockRegisterMockRecorder is the mock recorder for MockRegister.
type MockRegisterMockRecorder struct {
	mock *MockRegister
}

// NewMockRegister creates a new mock instance.
func NewMockRegister(ctrl *gomock.Controller) *MockRegister {
	mock := &MockRegister{ctrl: ctrl}
	mock.recorder = &MockRegisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegister) EXPECT() *MockRegisterMockRecorder {
	return m.recorder
}

// GenerateRoot mocks base method.
func (m *MockRegister) GenerateRoot(subscribeBlocks *SubscribeBlocks) (hash.Hash, *map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRoot", subscribeBlocks)
	ret0, _ := ret[0].(hash.Hash)
	ret1, _ := ret[1].(*map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateRoot indicates an expected call of GenerateRoot.
func (mr *MockRegisterMockRecorder) GenerateRoot(subscribeBlocks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRoot", reflect.TypeOf((*MockRegister)(nil).GenerateRoot), subscribeBlocks)
}

// GetLotteryWinners mocks base method.
func (m *MockRegister) GetLotteryWinners(peer string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLotteryWinners", peer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLotteryWinners indicates an expected call of GetLotteryWinners.
func (mr *MockRegisterMockRecorder) GetLotteryWinners(peer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLotteryWinners", reflect.TypeOf((*MockRegister)(nil).GetLotteryWinners), peer)
}

// IsClosed mocks base method.
func (m *MockRegister) IsClosed() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockRegisterMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockRegister)(nil).IsClosed))
}

// PickWinner mocks base method.
func (m *MockRegister) PickWinner() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PickWinner")
	ret0, _ := ret[0].(error)
	return ret0
}

// PickWinner indicates an expected call of PickWinner.
func (mr *MockRegisterMockRecorder) PickWinner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PickWinner", reflect.TypeOf((*MockRegister)(nil).PickWinner))
}
