// Code generated by MockGen. DO NOT EDIT.
// Source: event.go

// Package mock_event is a generated GoMock package.
package mock_event

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	event "github.com/modernice/goes/event"
)

// MockEvent is a mock of Event interface.
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
}

// MockEventMockRecorder is the mock recorder for MockEvent.
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance.
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// AggregateID mocks base method.
func (m *MockEvent) AggregateID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// AggregateID indicates an expected call of AggregateID.
func (mr *MockEventMockRecorder) AggregateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateID", reflect.TypeOf((*MockEvent)(nil).AggregateID))
}

// AggregateName mocks base method.
func (m *MockEvent) AggregateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AggregateName indicates an expected call of AggregateName.
func (mr *MockEventMockRecorder) AggregateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateName", reflect.TypeOf((*MockEvent)(nil).AggregateName))
}

// AggregateVersion mocks base method.
func (m *MockEvent) AggregateVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// AggregateVersion indicates an expected call of AggregateVersion.
func (mr *MockEventMockRecorder) AggregateVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateVersion", reflect.TypeOf((*MockEvent)(nil).AggregateVersion))
}

// Data mocks base method.
func (m *MockEvent) Data() event.Data {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(event.Data)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockEventMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockEvent)(nil).Data))
}

// ID mocks base method.
func (m *MockEvent) ID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockEventMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockEvent)(nil).ID))
}

// Name mocks base method.
func (m *MockEvent) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockEventMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockEvent)(nil).Name))
}

// Time mocks base method.
func (m *MockEvent) Time() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockEventMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockEvent)(nil).Time))
}

// MockData is a mock of Data interface.
type MockData struct {
	ctrl     *gomock.Controller
	recorder *MockDataMockRecorder
}

// MockDataMockRecorder is the mock recorder for MockData.
type MockDataMockRecorder struct {
	mock *MockData
}

// NewMockData creates a new mock instance.
func NewMockData(ctrl *gomock.Controller) *MockData {
	mock := &MockData{ctrl: ctrl}
	mock.recorder = &MockDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockData) EXPECT() *MockDataMockRecorder {
	return m.recorder
}
