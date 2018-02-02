// Automatically generated by MockGen. DO NOT EDIT!
// Source: api/iface/client.go

package mock

import (
	gomock "github.com/golang/mock/gomock"
	options "github.com/wata727/herogate/api/options"
	log "github.com/wata727/herogate/log"
)

// Mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockClientInterfaceRecorder
}

// Recorder for MockClientInterface (not exported)
type _MockClientInterfaceRecorder struct {
	mock *MockClientInterface
}

func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &_MockClientInterfaceRecorder{mock}
	return mock
}

func (_m *MockClientInterface) EXPECT() *_MockClientInterfaceRecorder {
	return _m.recorder
}

func (_m *MockClientInterface) DescribeLogs(appName string, options *options.DescribeLogs) []*log.Log {
	ret := _m.ctrl.Call(_m, "DescribeLogs", appName, options)
	ret0, _ := ret[0].([]*log.Log)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) DescribeLogs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogs", arg0, arg1)
}
