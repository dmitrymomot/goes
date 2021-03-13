// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/modernice/goes/command (interfaces: Command,Encoder,Bus,Registry)

// Package mock_command is a generated GoMock package.
package mock_command

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	command "github.com/modernice/goes/command"
)

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// AggregateID mocks base method.
func (m *MockCommand) AggregateID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// AggregateID indicates an expected call of AggregateID.
func (mr *MockCommandMockRecorder) AggregateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateID", reflect.TypeOf((*MockCommand)(nil).AggregateID))
}

// AggregateName mocks base method.
func (m *MockCommand) AggregateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AggregateName indicates an expected call of AggregateName.
func (mr *MockCommandMockRecorder) AggregateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateName", reflect.TypeOf((*MockCommand)(nil).AggregateName))
}

// ID mocks base method.
func (m *MockCommand) ID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockCommandMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockCommand)(nil).ID))
}

// Name mocks base method.
func (m *MockCommand) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCommandMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCommand)(nil).Name))
}

// Payload mocks base method.
func (m *MockCommand) Payload() command.Payload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload")
	ret0, _ := ret[0].(command.Payload)
	return ret0
}

// Payload indicates an expected call of Payload.
func (mr *MockCommandMockRecorder) Payload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockCommand)(nil).Payload))
}

// MockEncoder is a mock of Encoder interface.
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockEncoder) Decode(arg0 string, arg1 io.Reader) (command.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0, arg1)
	ret0, _ := ret[0].(command.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockEncoderMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockEncoder)(nil).Decode), arg0, arg1)
}

// Encode mocks base method.
func (m *MockEncoder) Encode(arg0 io.Writer, arg1 string, arg2 command.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderMockRecorder) Encode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0, arg1, arg2)
}

// MockBus is a mock of Bus interface.
type MockBus struct {
	ctrl     *gomock.Controller
	recorder *MockBusMockRecorder
}

// MockBusMockRecorder is the mock recorder for MockBus.
type MockBusMockRecorder struct {
	mock *MockBus
}

// NewMockBus creates a new mock instance.
func NewMockBus(ctrl *gomock.Controller) *MockBus {
	mock := &MockBus{ctrl: ctrl}
	mock.recorder = &MockBusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBus) EXPECT() *MockBusMockRecorder {
	return m.recorder
}

// Dispatch mocks base method.
func (m *MockBus) Dispatch(arg0 context.Context, arg1 command.Command, arg2 ...command.DispatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dispatch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockBusMockRecorder) Dispatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockBus)(nil).Dispatch), varargs...)
}

// Subscribe mocks base method.
func (m *MockBus) Subscribe(arg0 context.Context, arg1 ...string) (<-chan command.Context, <-chan error, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(<-chan command.Context)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBusMockRecorder) Subscribe(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBus)(nil).Subscribe), varargs...)
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockRegistry) Decode(arg0 string, arg1 io.Reader) (command.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0, arg1)
	ret0, _ := ret[0].(command.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockRegistryMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockRegistry)(nil).Decode), arg0, arg1)
}

// Encode mocks base method.
func (m *MockRegistry) Encode(arg0 io.Writer, arg1 string, arg2 command.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode.
func (mr *MockRegistryMockRecorder) Encode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockRegistry)(nil).Encode), arg0, arg1, arg2)
}

// New mocks base method.
func (m *MockRegistry) New(arg0 string) (command.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(command.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockRegistryMockRecorder) New(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockRegistry)(nil).New), arg0)
}

// Register mocks base method.
func (m *MockRegistry) Register(arg0 string, arg1 func() command.Payload) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0, arg1)
}

// Register indicates an expected call of Register.
func (mr *MockRegistryMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistry)(nil).Register), arg0, arg1)
}
