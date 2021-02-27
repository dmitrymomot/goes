// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_aggregate is a generated GoMock package.
package mock_aggregate

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	aggregate "github.com/modernice/goes/aggregate"
	version "github.com/modernice/goes/event/query/version"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockRepository) Save(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockRepositoryMockRecorder) Save(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, a)
}

// Fetch mocks base method
func (m *MockRepository) Fetch(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (mr *MockRepositoryMockRecorder) Fetch(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockRepository)(nil).Fetch), ctx, a)
}

// FetchVersion mocks base method
func (m *MockRepository) FetchVersion(ctx context.Context, a aggregate.Aggregate, v int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchVersion", ctx, a, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchVersion indicates an expected call of FetchVersion
func (mr *MockRepositoryMockRecorder) FetchVersion(ctx, a, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchVersion", reflect.TypeOf((*MockRepository)(nil).FetchVersion), ctx, a, v)
}

// Query mocks base method
func (m *MockRepository) Query(ctx context.Context, q aggregate.Query) (aggregate.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, q)
	ret0, _ := ret[0].(aggregate.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockRepositoryMockRecorder) Query(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRepository)(nil).Query), ctx, q)
}

// Delete mocks base method
func (m *MockRepository) Delete(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, a)
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Names mocks base method
func (m *MockQuery) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names
func (mr *MockQueryMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockQuery)(nil).Names))
}

// IDs mocks base method
func (m *MockQuery) IDs() []uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDs")
	ret0, _ := ret[0].([]uuid.UUID)
	return ret0
}

// IDs indicates an expected call of IDs
func (mr *MockQueryMockRecorder) IDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDs", reflect.TypeOf((*MockQuery)(nil).IDs))
}

// Versions mocks base method
func (m *MockQuery) Versions() version.Constraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Versions")
	ret0, _ := ret[0].(version.Constraints)
	return ret0
}

// Versions indicates an expected call of Versions
func (mr *MockQueryMockRecorder) Versions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Versions", reflect.TypeOf((*MockQuery)(nil).Versions))
}

// Sortings mocks base method
func (m *MockQuery) Sortings() []aggregate.SortOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sortings")
	ret0, _ := ret[0].([]aggregate.SortOptions)
	return ret0
}

// Sortings indicates an expected call of Sortings
func (mr *MockQueryMockRecorder) Sortings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sortings", reflect.TypeOf((*MockQuery)(nil).Sortings))
}

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockStream) Next(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockStreamMockRecorder) Next(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStream)(nil).Next), arg0)
}

// Current mocks base method
func (m *MockStream) Current() (string, uuid.UUID) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(uuid.UUID)
	return ret0, ret1
}

// Current indicates an expected call of Current
func (mr *MockStreamMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockStream)(nil).Current))
}

// Apply mocks base method
func (m *MockStream) Apply(arg0 aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockStreamMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockStream)(nil).Apply), arg0)
}

// Err mocks base method
func (m *MockStream) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockStreamMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockStream)(nil).Err))
}

// Close mocks base method
func (m *MockStream) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStreamMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close), arg0)
}
