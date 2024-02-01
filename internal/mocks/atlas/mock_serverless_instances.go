// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store/atlas (interfaces: ServerlessInstanceLister,ServerlessInstanceDescriber)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	atlas "github.com/andreaangiolillo/mongocli-test/internal/store/atlas"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockServerlessInstanceLister is a mock of ServerlessInstanceLister interface.
type MockServerlessInstanceLister struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceListerMockRecorder
}

// MockServerlessInstanceListerMockRecorder is the mock recorder for MockServerlessInstanceLister.
type MockServerlessInstanceListerMockRecorder struct {
	mock *MockServerlessInstanceLister
}

// NewMockServerlessInstanceLister creates a new mock instance.
func NewMockServerlessInstanceLister(ctrl *gomock.Controller) *MockServerlessInstanceLister {
	mock := &MockServerlessInstanceLister{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceLister) EXPECT() *MockServerlessInstanceListerMockRecorder {
	return m.recorder
}

// ServerlessInstances mocks base method.
func (m *MockServerlessInstanceLister) ServerlessInstances(arg0 string, arg1 *atlas.ListOptions) (*admin.PaginatedServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerlessInstances", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerlessInstances indicates an expected call of ServerlessInstances.
func (mr *MockServerlessInstanceListerMockRecorder) ServerlessInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerlessInstances", reflect.TypeOf((*MockServerlessInstanceLister)(nil).ServerlessInstances), arg0, arg1)
}

// MockServerlessInstanceDescriber is a mock of ServerlessInstanceDescriber interface.
type MockServerlessInstanceDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessInstanceDescriberMockRecorder
}

// MockServerlessInstanceDescriberMockRecorder is the mock recorder for MockServerlessInstanceDescriber.
type MockServerlessInstanceDescriberMockRecorder struct {
	mock *MockServerlessInstanceDescriber
}

// NewMockServerlessInstanceDescriber creates a new mock instance.
func NewMockServerlessInstanceDescriber(ctrl *gomock.Controller) *MockServerlessInstanceDescriber {
	mock := &MockServerlessInstanceDescriber{ctrl: ctrl}
	mock.recorder = &MockServerlessInstanceDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessInstanceDescriber) EXPECT() *MockServerlessInstanceDescriberMockRecorder {
	return m.recorder
}

// GetServerlessInstance mocks base method.
func (m *MockServerlessInstanceDescriber) GetServerlessInstance(arg0, arg1 string) (*admin.ServerlessInstanceDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessInstance", arg0, arg1)
	ret0, _ := ret[0].(*admin.ServerlessInstanceDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessInstance indicates an expected call of GetServerlessInstance.
func (mr *MockServerlessInstanceDescriberMockRecorder) GetServerlessInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessInstance", reflect.TypeOf((*MockServerlessInstanceDescriber)(nil).GetServerlessInstance), arg0, arg1)
}
