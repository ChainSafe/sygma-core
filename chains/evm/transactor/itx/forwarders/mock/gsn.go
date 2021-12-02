// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/transactor/itx/forwarders/gsn.go

// Package mock_forwarder is a generated GoMock package.
package mock_forwarder

import (
	big "math/big"
	reflect "reflect"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockForwarder is a mock of Forwarder interface.
type MockForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockForwarderMockRecorder
}

// MockForwarderMockRecorder is the mock recorder for MockForwarder.
type MockForwarderMockRecorder struct {
	mock *MockForwarder
}

// NewMockForwarder creates a new mock instance.
func NewMockForwarder(ctrl *gomock.Controller) *MockForwarder {
	mock := &MockForwarder{ctrl: ctrl}
	mock.recorder = &MockForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForwarder) EXPECT() *MockForwarderMockRecorder {
	return m.recorder
}

// ABI mocks base method.
func (m *MockForwarder) ABI() *abi.ABI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABI")
	ret0, _ := ret[0].(*abi.ABI)
	return ret0
}

// ABI indicates an expected call of ABI.
func (mr *MockForwarderMockRecorder) ABI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABI", reflect.TypeOf((*MockForwarder)(nil).ABI))
}

// Address mocks base method.
func (m *MockForwarder) Address() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockForwarderMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockForwarder)(nil).Address))
}

// GetNonce mocks base method.
func (m *MockForwarder) GetNonce(from common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", from)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockForwarderMockRecorder) GetNonce(from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockForwarder)(nil).GetNonce), from)
}