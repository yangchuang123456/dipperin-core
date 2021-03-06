// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/third-party/life/mem-manage (interfaces: SlabNeedInterface)

// Package mem_manage is a generated GoMock package.
package mem_manage

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSlabNeedInterface is a mock of SlabNeedInterface interface
type MockSlabNeedInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSlabNeedInterfaceMockRecorder
}

// MockSlabNeedInterfaceMockRecorder is the mock recorder for MockSlabNeedInterface
type MockSlabNeedInterfaceMockRecorder struct {
	mock *MockSlabNeedInterface
}

// NewMockSlabNeedInterface creates a new mock instance
func NewMockSlabNeedInterface(ctrl *gomock.Controller) *MockSlabNeedInterface {
	mock := &MockSlabNeedInterface{ctrl: ctrl}
	mock.recorder = &MockSlabNeedInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlabNeedInterface) EXPECT() *MockSlabNeedInterfaceMockRecorder {
	return m.recorder
}

// Free mocks base method
func (m *MockSlabNeedInterface) Free(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Free", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Free indicates an expected call of Free
func (mr *MockSlabNeedInterfaceMockRecorder) Free(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockSlabNeedInterface)(nil).Free), arg0)
}

// Malloc mocks base method
func (m *MockSlabNeedInterface) Malloc(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Malloc", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Malloc indicates an expected call of Malloc
func (mr *MockSlabNeedInterfaceMockRecorder) Malloc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Malloc", reflect.TypeOf((*MockSlabNeedInterface)(nil).Malloc), arg0)
}
