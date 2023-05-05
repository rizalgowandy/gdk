// Code generated by MockGen. DO NOT EDIT.
// Source: natsx.go

// Package natsx is a generated GoMock package.
package natsx

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockSubscriberItf is a mock of SubscriberItf interface.
type MockSubscriberItf struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberItfMockRecorder
}

// MockSubscriberItfMockRecorder is the mock recorder for MockSubscriberItf.
type MockSubscriberItfMockRecorder struct {
	mock *MockSubscriberItf
}

// NewMockSubscriberItf creates a new mock instance.
func NewMockSubscriberItf(ctrl *gomock.Controller) *MockSubscriberItf {
	mock := &MockSubscriberItf{ctrl: ctrl}
	mock.recorder = &MockSubscriberItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriberItf) EXPECT() *MockSubscriberItfMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockSubscriberItf) Handle(ctx context.Context, message *nats.Msg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockSubscriberItfMockRecorder) Handle(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockSubscriberItf)(nil).Handle), ctx, message)
}

// MockPublisherItf is a mock of PublisherItf interface.
type MockPublisherItf struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherItfMockRecorder
}

// MockPublisherItfMockRecorder is the mock recorder for MockPublisherItf.
type MockPublisherItfMockRecorder struct {
	mock *MockPublisherItf
}

// NewMockPublisherItf creates a new mock instance.
func NewMockPublisherItf(ctrl *gomock.Controller) *MockPublisherItf {
	mock := &MockPublisherItf{ctrl: ctrl}
	mock.recorder = &MockPublisherItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisherItf) EXPECT() *MockPublisherItfMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockPublisherItf) Publish(subj string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", subj, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockPublisherItfMockRecorder) Publish(subj, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPublisherItf)(nil).Publish), subj, data)
}