// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store (interfaces: DataLakeLister,DataLakeDescriber,DataLakeCreator,DataLakeDeleter,DataLakeUpdater,DataLakeStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockDataLakeLister is a mock of DataLakeLister interface.
type MockDataLakeLister struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeListerMockRecorder
}

// MockDataLakeListerMockRecorder is the mock recorder for MockDataLakeLister.
type MockDataLakeListerMockRecorder struct {
	mock *MockDataLakeLister
}

// NewMockDataLakeLister creates a new mock instance.
func NewMockDataLakeLister(ctrl *gomock.Controller) *MockDataLakeLister {
	mock := &MockDataLakeLister{ctrl: ctrl}
	mock.recorder = &MockDataLakeListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeLister) EXPECT() *MockDataLakeListerMockRecorder {
	return m.recorder
}

// DataLakes mocks base method.
func (m *MockDataLakeLister) DataLakes(arg0 string) ([]mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakes", arg0)
	ret0, _ := ret[0].([]mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakes indicates an expected call of DataLakes.
func (mr *MockDataLakeListerMockRecorder) DataLakes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakes", reflect.TypeOf((*MockDataLakeLister)(nil).DataLakes), arg0)
}

// MockDataLakeDescriber is a mock of DataLakeDescriber interface.
type MockDataLakeDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeDescriberMockRecorder
}

// MockDataLakeDescriberMockRecorder is the mock recorder for MockDataLakeDescriber.
type MockDataLakeDescriberMockRecorder struct {
	mock *MockDataLakeDescriber
}

// NewMockDataLakeDescriber creates a new mock instance.
func NewMockDataLakeDescriber(ctrl *gomock.Controller) *MockDataLakeDescriber {
	mock := &MockDataLakeDescriber{ctrl: ctrl}
	mock.recorder = &MockDataLakeDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeDescriber) EXPECT() *MockDataLakeDescriberMockRecorder {
	return m.recorder
}

// DataLake mocks base method.
func (m *MockDataLakeDescriber) DataLake(arg0, arg1 string) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLake", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLake indicates an expected call of DataLake.
func (mr *MockDataLakeDescriberMockRecorder) DataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLake", reflect.TypeOf((*MockDataLakeDescriber)(nil).DataLake), arg0, arg1)
}

// MockDataLakeCreator is a mock of DataLakeCreator interface.
type MockDataLakeCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeCreatorMockRecorder
}

// MockDataLakeCreatorMockRecorder is the mock recorder for MockDataLakeCreator.
type MockDataLakeCreatorMockRecorder struct {
	mock *MockDataLakeCreator
}

// NewMockDataLakeCreator creates a new mock instance.
func NewMockDataLakeCreator(ctrl *gomock.Controller) *MockDataLakeCreator {
	mock := &MockDataLakeCreator{ctrl: ctrl}
	mock.recorder = &MockDataLakeCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeCreator) EXPECT() *MockDataLakeCreatorMockRecorder {
	return m.recorder
}

// CreateDataLake mocks base method.
func (m *MockDataLakeCreator) CreateDataLake(arg0 string, arg1 *mongodbatlas.DataLakeCreateRequest) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataLake", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataLake indicates an expected call of CreateDataLake.
func (mr *MockDataLakeCreatorMockRecorder) CreateDataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataLake", reflect.TypeOf((*MockDataLakeCreator)(nil).CreateDataLake), arg0, arg1)
}

// MockDataLakeDeleter is a mock of DataLakeDeleter interface.
type MockDataLakeDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeDeleterMockRecorder
}

// MockDataLakeDeleterMockRecorder is the mock recorder for MockDataLakeDeleter.
type MockDataLakeDeleterMockRecorder struct {
	mock *MockDataLakeDeleter
}

// NewMockDataLakeDeleter creates a new mock instance.
func NewMockDataLakeDeleter(ctrl *gomock.Controller) *MockDataLakeDeleter {
	mock := &MockDataLakeDeleter{ctrl: ctrl}
	mock.recorder = &MockDataLakeDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeDeleter) EXPECT() *MockDataLakeDeleterMockRecorder {
	return m.recorder
}

// DeleteDataLake mocks base method.
func (m *MockDataLakeDeleter) DeleteDataLake(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataLake", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataLake indicates an expected call of DeleteDataLake.
func (mr *MockDataLakeDeleterMockRecorder) DeleteDataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataLake", reflect.TypeOf((*MockDataLakeDeleter)(nil).DeleteDataLake), arg0, arg1)
}

// MockDataLakeUpdater is a mock of DataLakeUpdater interface.
type MockDataLakeUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeUpdaterMockRecorder
}

// MockDataLakeUpdaterMockRecorder is the mock recorder for MockDataLakeUpdater.
type MockDataLakeUpdaterMockRecorder struct {
	mock *MockDataLakeUpdater
}

// NewMockDataLakeUpdater creates a new mock instance.
func NewMockDataLakeUpdater(ctrl *gomock.Controller) *MockDataLakeUpdater {
	mock := &MockDataLakeUpdater{ctrl: ctrl}
	mock.recorder = &MockDataLakeUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeUpdater) EXPECT() *MockDataLakeUpdaterMockRecorder {
	return m.recorder
}

// UpdateDataLake mocks base method.
func (m *MockDataLakeUpdater) UpdateDataLake(arg0, arg1 string, arg2 *mongodbatlas.DataLakeUpdateRequest) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataLake", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataLake indicates an expected call of UpdateDataLake.
func (mr *MockDataLakeUpdaterMockRecorder) UpdateDataLake(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDataLake", reflect.TypeOf((*MockDataLakeUpdater)(nil).UpdateDataLake), arg0, arg1, arg2)
}

// MockDataLakeStore is a mock of DataLakeStore interface.
type MockDataLakeStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeStoreMockRecorder
}

// MockDataLakeStoreMockRecorder is the mock recorder for MockDataLakeStore.
type MockDataLakeStoreMockRecorder struct {
	mock *MockDataLakeStore
}

// NewMockDataLakeStore creates a new mock instance.
func NewMockDataLakeStore(ctrl *gomock.Controller) *MockDataLakeStore {
	mock := &MockDataLakeStore{ctrl: ctrl}
	mock.recorder = &MockDataLakeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataLakeStore) EXPECT() *MockDataLakeStoreMockRecorder {
	return m.recorder
}

// CreateDataLake mocks base method.
func (m *MockDataLakeStore) CreateDataLake(arg0 string, arg1 *mongodbatlas.DataLakeCreateRequest) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataLake", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataLake indicates an expected call of CreateDataLake.
func (mr *MockDataLakeStoreMockRecorder) CreateDataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataLake", reflect.TypeOf((*MockDataLakeStore)(nil).CreateDataLake), arg0, arg1)
}

// DataLake mocks base method.
func (m *MockDataLakeStore) DataLake(arg0, arg1 string) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLake", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLake indicates an expected call of DataLake.
func (mr *MockDataLakeStoreMockRecorder) DataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLake", reflect.TypeOf((*MockDataLakeStore)(nil).DataLake), arg0, arg1)
}

// DataLakes mocks base method.
func (m *MockDataLakeStore) DataLakes(arg0 string) ([]mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataLakes", arg0)
	ret0, _ := ret[0].([]mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataLakes indicates an expected call of DataLakes.
func (mr *MockDataLakeStoreMockRecorder) DataLakes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataLakes", reflect.TypeOf((*MockDataLakeStore)(nil).DataLakes), arg0)
}

// DeleteDataLake mocks base method.
func (m *MockDataLakeStore) DeleteDataLake(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataLake", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataLake indicates an expected call of DeleteDataLake.
func (mr *MockDataLakeStoreMockRecorder) DeleteDataLake(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataLake", reflect.TypeOf((*MockDataLakeStore)(nil).DeleteDataLake), arg0, arg1)
}

// UpdateDataLake mocks base method.
func (m *MockDataLakeStore) UpdateDataLake(arg0, arg1 string, arg2 *mongodbatlas.DataLakeUpdateRequest) (*mongodbatlas.DataLake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataLake", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.DataLake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataLake indicates an expected call of UpdateDataLake.
func (mr *MockDataLakeStoreMockRecorder) UpdateDataLake(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDataLake", reflect.TypeOf((*MockDataLakeStore)(nil).UpdateDataLake), arg0, arg1, arg2)
}
