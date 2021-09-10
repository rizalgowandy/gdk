// Code generated by MockGen. DO NOT EDIT.
// Source: httpx.go

// Package httpx is a generated GoMock package.
package httpx

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClientItf is a mock of ClientItf interface.
type MockClientItf struct {
	ctrl     *gomock.Controller
	recorder *MockClientItfMockRecorder
}

// MockClientItfMockRecorder is the mock recorder for MockClientItf.
type MockClientItfMockRecorder struct {
	mock *MockClientItf
}

// NewMockClientItf creates a new mock instance.
func NewMockClientItf(ctrl *gomock.Controller) *MockClientItf {
	mock := &MockClientItf{ctrl: ctrl}
	mock.recorder = &MockClientItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientItf) EXPECT() *MockClientItfMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockClientItf) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockClientItfMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClientItf)(nil).Do), req)
}

// MockReadCloserItf is a mock of ReadCloserItf interface.
type MockReadCloserItf struct {
	ctrl     *gomock.Controller
	recorder *MockReadCloserItfMockRecorder
}

// MockReadCloserItfMockRecorder is the mock recorder for MockReadCloserItf.
type MockReadCloserItfMockRecorder struct {
	mock *MockReadCloserItf
}

// NewMockReadCloserItf creates a new mock instance.
func NewMockReadCloserItf(ctrl *gomock.Controller) *MockReadCloserItf {
	mock := &MockReadCloserItf{ctrl: ctrl}
	mock.recorder = &MockReadCloserItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadCloserItf) EXPECT() *MockReadCloserItfMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReadCloserItf) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReadCloserItfMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadCloserItf)(nil).Close))
}

// Read mocks base method.
func (m *MockReadCloserItf) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReadCloserItfMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadCloserItf)(nil).Read), arg0)
}
