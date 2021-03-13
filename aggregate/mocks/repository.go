// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_aggregate is a generated GoMock package.
package mock_aggregate

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	aggregate "github.com/modernice/goes/aggregate"
	version "github.com/modernice/goes/event/query/version"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, a)
}

// Fetch mocks base method.
func (m *MockRepository) Fetch(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockRepositoryMockRecorder) Fetch(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockRepository)(nil).Fetch), ctx, a)
}

// FetchVersion mocks base method.
func (m *MockRepository) FetchVersion(ctx context.Context, a aggregate.Aggregate, v int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchVersion", ctx, a, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchVersion indicates an expected call of FetchVersion.
func (mr *MockRepositoryMockRecorder) FetchVersion(ctx, a, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchVersion", reflect.TypeOf((*MockRepository)(nil).FetchVersion), ctx, a, v)
}

// Query mocks base method.
func (m *MockRepository) Query(ctx context.Context, q aggregate.Query) (<-chan aggregate.History, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, q)
	ret0, _ := ret[0].(<-chan aggregate.History)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockRepositoryMockRecorder) Query(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRepository)(nil).Query), ctx, q)
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, a aggregate.Aggregate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, a)
}

// MockQuery is a mock of Query interface.
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance.
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// IDs mocks base method.
func (m *MockQuery) IDs() []uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDs")
	ret0, _ := ret[0].([]uuid.UUID)
	return ret0
}

// IDs indicates an expected call of IDs.
func (mr *MockQueryMockRecorder) IDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDs", reflect.TypeOf((*MockQuery)(nil).IDs))
}

// Names mocks base method.
func (m *MockQuery) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockQueryMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockQuery)(nil).Names))
}

// Sortings mocks base method.
func (m *MockQuery) Sortings() []aggregate.SortOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sortings")
	ret0, _ := ret[0].([]aggregate.SortOptions)
	return ret0
}

// Sortings indicates an expected call of Sortings.
func (mr *MockQueryMockRecorder) Sortings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sortings", reflect.TypeOf((*MockQuery)(nil).Sortings))
}

// Versions mocks base method.
func (m *MockQuery) Versions() version.Constraints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Versions")
	ret0, _ := ret[0].(version.Constraints)
	return ret0
}

// Versions indicates an expected call of Versions.
func (mr *MockQueryMockRecorder) Versions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Versions", reflect.TypeOf((*MockQuery)(nil).Versions))
}

// MockHistory is a mock of History interface.
type MockHistory struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryMockRecorder
}

// MockHistoryMockRecorder is the mock recorder for MockHistory.
type MockHistoryMockRecorder struct {
	mock *MockHistory
}

// NewMockHistory creates a new mock instance.
func NewMockHistory(ctrl *gomock.Controller) *MockHistory {
	mock := &MockHistory{ctrl: ctrl}
	mock.recorder = &MockHistoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistory) EXPECT() *MockHistoryMockRecorder {
	return m.recorder
}

// AggregateID mocks base method.
func (m *MockHistory) AggregateID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// AggregateID indicates an expected call of AggregateID.
func (mr *MockHistoryMockRecorder) AggregateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateID", reflect.TypeOf((*MockHistory)(nil).AggregateID))
}

// AggregateName mocks base method.
func (m *MockHistory) AggregateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AggregateName indicates an expected call of AggregateName.
func (mr *MockHistoryMockRecorder) AggregateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateName", reflect.TypeOf((*MockHistory)(nil).AggregateName))
}

// Apply mocks base method.
func (m *MockHistory) Apply(arg0 aggregate.Aggregate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Apply", arg0)
}

// Apply indicates an expected call of Apply.
func (mr *MockHistoryMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockHistory)(nil).Apply), arg0)
}
