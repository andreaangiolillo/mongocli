// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store/atlas (interfaces: MaintenanceWindowDescriber)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockMaintenanceWindowDescriber is a mock of MaintenanceWindowDescriber interface.
type MockMaintenanceWindowDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowDescriberMockRecorder
}

// MockMaintenanceWindowDescriberMockRecorder is the mock recorder for MockMaintenanceWindowDescriber.
type MockMaintenanceWindowDescriberMockRecorder struct {
	mock *MockMaintenanceWindowDescriber
}

// NewMockMaintenanceWindowDescriber creates a new mock instance.
func NewMockMaintenanceWindowDescriber(ctrl *gomock.Controller) *MockMaintenanceWindowDescriber {
	mock := &MockMaintenanceWindowDescriber{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceWindowDescriber) EXPECT() *MockMaintenanceWindowDescriberMockRecorder {
	return m.recorder
}

// MaintenanceWindow mocks base method.
func (m *MockMaintenanceWindowDescriber) MaintenanceWindow(arg0 string) (*admin.GroupMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*admin.GroupMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockMaintenanceWindowDescriberMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowDescriber)(nil).MaintenanceWindow), arg0)
}
