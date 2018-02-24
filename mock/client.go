// Code generated by MockGen. DO NOT EDIT.
// Source: iface/client.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	objects "github.com/wata727/herogate/api/objects"
	options "github.com/wata727/herogate/api/options"
	log "github.com/wata727/herogate/log"
	reflect "reflect"
)

// MockClientInterface is a mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// CreateApp mocks base method
func (m *MockClientInterface) CreateApp(appName string) *objects.App {
	ret := m.ctrl.Call(m, "CreateApp", appName)
	ret0, _ := ret[0].(*objects.App)
	return ret0
}

// CreateApp indicates an expected call of CreateApp
func (mr *MockClientInterfaceMockRecorder) CreateApp(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApp", reflect.TypeOf((*MockClientInterface)(nil).CreateApp), appName)
}

// GetAppCreationProgress mocks base method
func (m *MockClientInterface) GetAppCreationProgress(appName string) int {
	ret := m.ctrl.Call(m, "GetAppCreationProgress", appName)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAppCreationProgress indicates an expected call of GetAppCreationProgress
func (mr *MockClientInterfaceMockRecorder) GetAppCreationProgress(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppCreationProgress", reflect.TypeOf((*MockClientInterface)(nil).GetAppCreationProgress), appName)
}

// DescribeLogs mocks base method
func (m *MockClientInterface) DescribeLogs(appName string, options *options.DescribeLogs) ([]*log.Log, error) {
	ret := m.ctrl.Call(m, "DescribeLogs", appName, options)
	ret0, _ := ret[0].([]*log.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogs indicates an expected call of DescribeLogs
func (mr *MockClientInterfaceMockRecorder) DescribeLogs(appName, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogs", reflect.TypeOf((*MockClientInterface)(nil).DescribeLogs), appName, options)
}

// GetApp mocks base method
func (m *MockClientInterface) GetApp(appName string) (*objects.App, error) {
	ret := m.ctrl.Call(m, "GetApp", appName)
	ret0, _ := ret[0].(*objects.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp
func (mr *MockClientInterfaceMockRecorder) GetApp(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockClientInterface)(nil).GetApp), appName)
}

// GetTemplate mocks base method
func (m *MockClientInterface) GetTemplate(appName string) string {
	ret := m.ctrl.Call(m, "GetTemplate", appName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTemplate indicates an expected call of GetTemplate
func (mr *MockClientInterfaceMockRecorder) GetTemplate(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockClientInterface)(nil).GetTemplate), appName)
}

// DestroyApp mocks base method
func (m *MockClientInterface) DestroyApp(appName string) error {
	ret := m.ctrl.Call(m, "DestroyApp", appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyApp indicates an expected call of DestroyApp
func (mr *MockClientInterfaceMockRecorder) DestroyApp(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyApp", reflect.TypeOf((*MockClientInterface)(nil).DestroyApp), appName)
}

// GetAppDeletionProgress mocks base method
func (m *MockClientInterface) GetAppDeletionProgress(appName string) int {
	ret := m.ctrl.Call(m, "GetAppDeletionProgress", appName)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAppDeletionProgress indicates an expected call of GetAppDeletionProgress
func (mr *MockClientInterfaceMockRecorder) GetAppDeletionProgress(appName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppDeletionProgress", reflect.TypeOf((*MockClientInterface)(nil).GetAppDeletionProgress), appName)
}
