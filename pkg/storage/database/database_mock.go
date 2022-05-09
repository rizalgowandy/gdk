// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package database is a generated GoMock package.
package database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

// MockPostgreClientItf is a mock of PostgreClientItf interface.
type MockPostgreClientItf struct {
	ctrl     *gomock.Controller
	recorder *MockPostgreClientItfMockRecorder
}

// MockPostgreClientItfMockRecorder is the mock recorder for MockPostgreClientItf.
type MockPostgreClientItfMockRecorder struct {
	mock *MockPostgreClientItf
}

// NewMockPostgreClientItf creates a new mock instance.
func NewMockPostgreClientItf(ctrl *gomock.Controller) *MockPostgreClientItf {
	mock := &MockPostgreClientItf{ctrl: ctrl}
	mock.recorder = &MockPostgreClientItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgreClientItf) EXPECT() *MockPostgreClientItfMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPostgreClientItf) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockPostgreClientItfMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPostgreClientItf)(nil).Close))
}

// Get mocks base method.
func (m *MockPostgreClientItf) Get(ctx context.Context) (*pgxpool.Pool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(*pgxpool.Pool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostgreClientItfMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostgreClientItf)(nil).Get), ctx)
}

// GetReader mocks base method.
func (m *MockPostgreClientItf) GetReader(ctx context.Context) (*pgxpool.Pool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReader", ctx)
	ret0, _ := ret[0].(*pgxpool.Pool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReader indicates an expected call of GetReader.
func (mr *MockPostgreClientItfMockRecorder) GetReader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReader", reflect.TypeOf((*MockPostgreClientItf)(nil).GetReader), ctx)
}

// GetWriter mocks base method.
func (m *MockPostgreClientItf) GetWriter(ctx context.Context) (*pgxpool.Pool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWriter", ctx)
	ret0, _ := ret[0].(*pgxpool.Pool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWriter indicates an expected call of GetWriter.
func (mr *MockPostgreClientItfMockRecorder) GetWriter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriter", reflect.TypeOf((*MockPostgreClientItf)(nil).GetWriter), ctx)
}
