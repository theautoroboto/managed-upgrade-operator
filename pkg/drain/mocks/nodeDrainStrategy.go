// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/drain (interfaces: NodeDrainStrategy)

// Package mocks is a generated GoMock package.
package mocks

import (
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	drain "github.com/openshift/managed-upgrade-operator/pkg/drain"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockNodeDrainStrategy is a mock of NodeDrainStrategy interface
type MockNodeDrainStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockNodeDrainStrategyMockRecorder
}

// MockNodeDrainStrategyMockRecorder is the mock recorder for MockNodeDrainStrategy
type MockNodeDrainStrategyMockRecorder struct {
	mock *MockNodeDrainStrategy
}

// NewMockNodeDrainStrategy creates a new mock instance
func NewMockNodeDrainStrategy(ctrl *gomock.Controller) *MockNodeDrainStrategy {
	mock := &MockNodeDrainStrategy{ctrl: ctrl}
	mock.recorder = &MockNodeDrainStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeDrainStrategy) EXPECT() *MockNodeDrainStrategyMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockNodeDrainStrategy) Execute(arg0 *v1.Node, arg1 logr.Logger) ([]*drain.DrainStrategyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].([]*drain.DrainStrategyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockNodeDrainStrategyMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockNodeDrainStrategy)(nil).Execute), arg0, arg1)
}

// HasFailed mocks base method
func (m *MockNodeDrainStrategy) HasFailed(arg0 *v1.Node, arg1 logr.Logger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFailed", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasFailed indicates an expected call of HasFailed
func (mr *MockNodeDrainStrategyMockRecorder) HasFailed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFailed", reflect.TypeOf((*MockNodeDrainStrategy)(nil).HasFailed), arg0, arg1)
}
