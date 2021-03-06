// Code generated by MockGen. DO NOT EDIT.
// Source: bugreport.go

package bugreport

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExample is a mock of Example interface
type MockExample struct {
	ctrl     *gomock.Controller
	recorder *MockExampleMockRecorder
}

// MockExampleMockRecorder is the mock recorder for MockExample
type MockExampleMockRecorder struct {
	mock *MockExample
}

// NewMockExample creates a new mock instance
func NewMockExample(ctrl *gomock.Controller) *MockExample {
	mock := &MockExample{ctrl: ctrl}
	mock.recorder = &MockExampleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockExample) EXPECT() *MockExampleMockRecorder {
	return _m.recorder
}

// someMethod mocks base method
func (_m *MockExample) someMethod(_param0 string) string {
	ret := _m.ctrl.Call(_m, "someMethod", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

// someMethod indicates an expected call of someMethod
func (_mr *MockExampleMockRecorder) someMethod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "someMethod", reflect.TypeOf((*MockExample)(nil).someMethod), arg0)
}
