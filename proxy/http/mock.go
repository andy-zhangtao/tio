// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockdataLoader is a mock of dataLoader interface
type MockdataLoader struct {
	ctrl     *gomock.Controller
	recorder *MockdataLoaderMockRecorder
}

// MockdataLoaderMockRecorder is the mock recorder for MockdataLoader
type MockdataLoaderMockRecorder struct {
	mock *MockdataLoader
}

// NewMockdataLoader creates a new mock instance
func NewMockdataLoader(ctrl *gomock.Controller) *MockdataLoader {
	mock := &MockdataLoader{ctrl: ctrl}
	mock.recorder = &MockdataLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdataLoader) EXPECT() *MockdataLoaderMockRecorder {
	return m.recorder
}

// LoadInjectData mocks base method
func (m *MockdataLoader) LoadInjectData() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadInjectData")
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadInjectData indicates an expected call of LoadInjectData
func (mr *MockdataLoaderMockRecorder) LoadInjectData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadInjectData", reflect.TypeOf((*MockdataLoader)(nil).LoadInjectData))
}

// Scala mocks base method
func (m *MockdataLoader) Scala(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scala", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scala indicates an expected call of Scala
func (mr *MockdataLoaderMockRecorder) Scala(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scala", reflect.TypeOf((*MockdataLoader)(nil).Scala), arg0)
}

// Wait mocks base method
func (m *MockdataLoader) Wait(arg0 string) (service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait
func (mr *MockdataLoaderMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockdataLoader)(nil).Wait), arg0)
}

// Done mocks base method
func (m *MockdataLoader) Done(arg0 service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockdataLoaderMockRecorder) Done(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockdataLoader)(nil).Done), arg0)
}

// IsValidInject mocks base method
func (m *MockdataLoader) IsValidInject(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidInject", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidInject indicates an expected call of IsValidInject
func (mr *MockdataLoaderMockRecorder) IsValidInject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidInject", reflect.TypeOf((*MockdataLoader)(nil).IsValidInject), arg0)
}

// GetServiceName mocks base method
func (m *MockdataLoader) GetServiceName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceName indicates an expected call of GetServiceName
func (mr *MockdataLoaderMockRecorder) GetServiceName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceName", reflect.TypeOf((*MockdataLoader)(nil).GetServiceName), arg0)
}

// Transfer mocks base method
func (m *MockdataLoader) Transfer(arg0 string, arg1 http.ResponseWriter, arg2 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Transfer", arg0, arg1, arg2)
}

// Transfer indicates an expected call of Transfer
func (mr *MockdataLoaderMockRecorder) Transfer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockdataLoader)(nil).Transfer), arg0, arg1, arg2)
}
