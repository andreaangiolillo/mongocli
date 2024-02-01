// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store (interfaces: LiveMigrationCreator,LiveMigrationDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockLiveMigrationCreator is a mock of LiveMigrationCreator interface.
type MockLiveMigrationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationCreatorMockRecorder
}

// MockLiveMigrationCreatorMockRecorder is the mock recorder for MockLiveMigrationCreator.
type MockLiveMigrationCreatorMockRecorder struct {
	mock *MockLiveMigrationCreator
}

// NewMockLiveMigrationCreator creates a new mock instance.
func NewMockLiveMigrationCreator(ctrl *gomock.Controller) *MockLiveMigrationCreator {
	mock := &MockLiveMigrationCreator{ctrl: ctrl}
	mock.recorder = &MockLiveMigrationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMigrationCreator) EXPECT() *MockLiveMigrationCreatorMockRecorder {
	return m.recorder
}

// LiveMigrationCreate mocks base method.
func (m *MockLiveMigrationCreator) LiveMigrationCreate(arg0 string, arg1 *admin.LiveMigrationRequest) (*admin.LiveMigrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LiveMigrationCreate", arg0, arg1)
	ret0, _ := ret[0].(*admin.LiveMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LiveMigrationCreate indicates an expected call of LiveMigrationCreate.
func (mr *MockLiveMigrationCreatorMockRecorder) LiveMigrationCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiveMigrationCreate", reflect.TypeOf((*MockLiveMigrationCreator)(nil).LiveMigrationCreate), arg0, arg1)
}

// MockLiveMigrationDescriber is a mock of LiveMigrationDescriber interface.
type MockLiveMigrationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMigrationDescriberMockRecorder
}

// MockLiveMigrationDescriberMockRecorder is the mock recorder for MockLiveMigrationDescriber.
type MockLiveMigrationDescriberMockRecorder struct {
	mock *MockLiveMigrationDescriber
}

// NewMockLiveMigrationDescriber creates a new mock instance.
func NewMockLiveMigrationDescriber(ctrl *gomock.Controller) *MockLiveMigrationDescriber {
	mock := &MockLiveMigrationDescriber{ctrl: ctrl}
	mock.recorder = &MockLiveMigrationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiveMigrationDescriber) EXPECT() *MockLiveMigrationDescriberMockRecorder {
	return m.recorder
}

// LiveMigrationDescribe mocks base method.
func (m *MockLiveMigrationDescriber) LiveMigrationDescribe(arg0, arg1 string) (*admin.LiveMigrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LiveMigrationDescribe", arg0, arg1)
	ret0, _ := ret[0].(*admin.LiveMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LiveMigrationDescribe indicates an expected call of LiveMigrationDescribe.
func (mr *MockLiveMigrationDescriberMockRecorder) LiveMigrationDescribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiveMigrationDescribe", reflect.TypeOf((*MockLiveMigrationDescriber)(nil).LiveMigrationDescribe), arg0, arg1)
}
