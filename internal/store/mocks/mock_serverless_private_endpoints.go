// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store/atlas (interfaces: ServerlessPrivateEndpointsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockServerlessPrivateEndpointsLister is a mock of ServerlessPrivateEndpointsLister interface.
type MockServerlessPrivateEndpointsLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessPrivateEndpointsListerMockRecorder
}

// MockServerlessPrivateEndpointsListerMockRecorder is the mock recorder for MockServerlessPrivateEndpointsLister.
type MockServerlessPrivateEndpointsListerMockRecorder struct {
	mock *MockServerlessPrivateEndpointsLister
}

// NewMockServerlessPrivateEndpointsLister creates a new mock instance.
func NewMockServerlessPrivateEndpointsLister(ctrl *gomock.Controller) *MockServerlessPrivateEndpointsLister {
	mock := &MockServerlessPrivateEndpointsLister{ctrl: ctrl}
	mock.recorder = &MockServerlessPrivateEndpointsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessPrivateEndpointsLister) EXPECT() *MockServerlessPrivateEndpointsListerMockRecorder {
	return m.recorder
}

// ServerlessPrivateEndpoints mocks base method.
func (m *MockServerlessPrivateEndpointsLister) ServerlessPrivateEndpoints(arg0, arg1 string) ([]admin.ServerlessTenantEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessPrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.ServerlessTenantEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessPrivateEndpoints indicates an expected call of ServerlessPrivateEndpoints.
func (mr *MockServerlessPrivateEndpointsListerMockRecorder) ServerlessPrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessPrivateEndpoints", reflect.TypeOf((*MockServerlessPrivateEndpointsLister)(nil).ServerlessPrivateEndpoints), arg0, arg1)
}
