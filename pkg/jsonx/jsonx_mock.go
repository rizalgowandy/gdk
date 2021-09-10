// Code generated by MockGen. DO NOT EDIT.
// Source: jsonx.go

// Package jsonx is a generated GoMock package.
package jsonx

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOperatorItf is a mock of OperatorItf interface.
type MockOperatorItf struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorItfMockRecorder
}

// MockOperatorItfMockRecorder is the mock recorder for MockOperatorItf.
type MockOperatorItfMockRecorder struct {
	mock *MockOperatorItf
}

// NewMockOperatorItf creates a new mock instance.
func NewMockOperatorItf(ctrl *gomock.Controller) *MockOperatorItf {
	mock := &MockOperatorItf{ctrl: ctrl}
	mock.recorder = &MockOperatorItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorItf) EXPECT() *MockOperatorItfMockRecorder {
	return m.recorder
}

// Marshal mocks base method.
func (m *MockOperatorItf) Marshal(v interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal", v)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal.
func (mr *MockOperatorItfMockRecorder) Marshal(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockOperatorItf)(nil).Marshal), v)
}

// NewDecoder mocks base method.
func (m *MockOperatorItf) NewDecoder(r io.Reader) DecoderItf {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDecoder", r)
	ret0, _ := ret[0].(DecoderItf)
	return ret0
}

// NewDecoder indicates an expected call of NewDecoder.
func (mr *MockOperatorItfMockRecorder) NewDecoder(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDecoder", reflect.TypeOf((*MockOperatorItf)(nil).NewDecoder), r)
}

// NewEncoder mocks base method.
func (m *MockOperatorItf) NewEncoder(w io.Writer) EncoderItf {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEncoder", w)
	ret0, _ := ret[0].(EncoderItf)
	return ret0
}

// NewEncoder indicates an expected call of NewEncoder.
func (mr *MockOperatorItfMockRecorder) NewEncoder(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEncoder", reflect.TypeOf((*MockOperatorItf)(nil).NewEncoder), w)
}

// Unmarshal mocks base method.
func (m *MockOperatorItf) Unmarshal(data []byte, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockOperatorItfMockRecorder) Unmarshal(data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockOperatorItf)(nil).Unmarshal), data, v)
}

// MockEncoderItf is a mock of EncoderItf interface.
type MockEncoderItf struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderItfMockRecorder
}

// MockEncoderItfMockRecorder is the mock recorder for MockEncoderItf.
type MockEncoderItfMockRecorder struct {
	mock *MockEncoderItf
}

// NewMockEncoderItf creates a new mock instance.
func NewMockEncoderItf(ctrl *gomock.Controller) *MockEncoderItf {
	mock := &MockEncoderItf{ctrl: ctrl}
	mock.recorder = &MockEncoderItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoderItf) EXPECT() *MockEncoderItfMockRecorder {
	return m.recorder
}

// Encode mocks base method.
func (m *MockEncoderItf) Encode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderItfMockRecorder) Encode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoderItf)(nil).Encode), v)
}

// MockDecoderItf is a mock of DecoderItf interface.
type MockDecoderItf struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderItfMockRecorder
}

// MockDecoderItfMockRecorder is the mock recorder for MockDecoderItf.
type MockDecoderItfMockRecorder struct {
	mock *MockDecoderItf
}

// NewMockDecoderItf creates a new mock instance.
func NewMockDecoderItf(ctrl *gomock.Controller) *MockDecoderItf {
	mock := &MockDecoderItf{ctrl: ctrl}
	mock.recorder = &MockDecoderItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoderItf) EXPECT() *MockDecoderItfMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoderItf) Decode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderItfMockRecorder) Decode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoderItf)(nil).Decode), v)
}
