// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store/atlas (interfaces: GlobalClusterDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockGlobalClusterDescriber is a mock of GlobalClusterDescriber interface.
type MockGlobalClusterDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalClusterDescriberMockRecorder
}

// MockGlobalClusterDescriberMockRecorder is the mock recorder for MockGlobalClusterDescriber.
type MockGlobalClusterDescriberMockRecorder struct {
	mock *MockGlobalClusterDescriber
}

// NewMockGlobalClusterDescriber creates a new mock instance.
func NewMockGlobalClusterDescriber(ctrl *gomock.Controller) *MockGlobalClusterDescriber {
	mock := &MockGlobalClusterDescriber{ctrl: ctrl}
	mock.recorder = &MockGlobalClusterDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalClusterDescriber) EXPECT() *MockGlobalClusterDescriberMockRecorder {
	return m.recorder
}

// GlobalCluster mocks base method.
func (m *MockGlobalClusterDescriber) GlobalCluster(arg0, arg1 string) (*admin.GeoSharding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalCluster", arg0, arg1)
	ret0, _ := ret[0].(*admin.GeoSharding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GlobalCluster indicates an expected call of GlobalCluster.
func (mr *MockGlobalClusterDescriberMockRecorder) GlobalCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalCluster", reflect.TypeOf((*MockGlobalClusterDescriber)(nil).GlobalCluster), arg0, arg1)
}
