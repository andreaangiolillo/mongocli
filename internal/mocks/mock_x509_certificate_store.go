// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store (interfaces: X509CertificateConfDescriber,X509CertificateConfSaver,X509CertificateConfDisabler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockX509CertificateConfDescriber is a mock of X509CertificateConfDescriber interface.
type MockX509CertificateConfDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateConfDescriberMockRecorder
}

// MockX509CertificateConfDescriberMockRecorder is the mock recorder for MockX509CertificateConfDescriber.
type MockX509CertificateConfDescriberMockRecorder struct {
	mock *MockX509CertificateConfDescriber
}

// NewMockX509CertificateConfDescriber creates a new mock instance.
func NewMockX509CertificateConfDescriber(ctrl *gomock.Controller) *MockX509CertificateConfDescriber {
	mock := &MockX509CertificateConfDescriber{ctrl: ctrl}
	mock.recorder = &MockX509CertificateConfDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockX509CertificateConfDescriber) EXPECT() *MockX509CertificateConfDescriberMockRecorder {
	return m.recorder
}

// X509Configuration mocks base method.
func (m *MockX509CertificateConfDescriber) X509Configuration(arg0 string) (*admin.UserSecurity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "X509Configuration", arg0)
	ret0, _ := ret[0].(*admin.UserSecurity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// X509Configuration indicates an expected call of X509Configuration.
func (mr *MockX509CertificateConfDescriberMockRecorder) X509Configuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "X509Configuration", reflect.TypeOf((*MockX509CertificateConfDescriber)(nil).X509Configuration), arg0)
}

// MockX509CertificateConfSaver is a mock of X509CertificateConfSaver interface.
type MockX509CertificateConfSaver struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateConfSaverMockRecorder
}

// MockX509CertificateConfSaverMockRecorder is the mock recorder for MockX509CertificateConfSaver.
type MockX509CertificateConfSaverMockRecorder struct {
	mock *MockX509CertificateConfSaver
}

// NewMockX509CertificateConfSaver creates a new mock instance.
func NewMockX509CertificateConfSaver(ctrl *gomock.Controller) *MockX509CertificateConfSaver {
	mock := &MockX509CertificateConfSaver{ctrl: ctrl}
	mock.recorder = &MockX509CertificateConfSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockX509CertificateConfSaver) EXPECT() *MockX509CertificateConfSaverMockRecorder {
	return m.recorder
}

// SaveX509Configuration mocks base method.
func (m *MockX509CertificateConfSaver) SaveX509Configuration(arg0, arg1 string) (*admin.UserSecurity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveX509Configuration", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserSecurity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveX509Configuration indicates an expected call of SaveX509Configuration.
func (mr *MockX509CertificateConfSaverMockRecorder) SaveX509Configuration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveX509Configuration", reflect.TypeOf((*MockX509CertificateConfSaver)(nil).SaveX509Configuration), arg0, arg1)
}

// MockX509CertificateConfDisabler is a mock of X509CertificateConfDisabler interface.
type MockX509CertificateConfDisabler struct {
	ctrl     *gomock.Controller
	recorder *MockX509CertificateConfDisablerMockRecorder
}

// MockX509CertificateConfDisablerMockRecorder is the mock recorder for MockX509CertificateConfDisabler.
type MockX509CertificateConfDisablerMockRecorder struct {
	mock *MockX509CertificateConfDisabler
}

// NewMockX509CertificateConfDisabler creates a new mock instance.
func NewMockX509CertificateConfDisabler(ctrl *gomock.Controller) *MockX509CertificateConfDisabler {
	mock := &MockX509CertificateConfDisabler{ctrl: ctrl}
	mock.recorder = &MockX509CertificateConfDisablerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockX509CertificateConfDisabler) EXPECT() *MockX509CertificateConfDisablerMockRecorder {
	return m.recorder
}

// DisableX509Configuration mocks base method.
func (m *MockX509CertificateConfDisabler) DisableX509Configuration(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableX509Configuration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableX509Configuration indicates an expected call of DisableX509Configuration.
func (mr *MockX509CertificateConfDisablerMockRecorder) DisableX509Configuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableX509Configuration", reflect.TypeOf((*MockX509CertificateConfDisabler)(nil).DisableX509Configuration), arg0)
}
