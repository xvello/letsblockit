// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/db/querier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	pgtype "github.com/jackc/pgtype"
	db "github.com/xvello/letsblockit/src/db"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CountInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) CountInstanceForUserAndFilter(ctx context.Context, arg db.CountInstanceForUserAndFilterParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountInstanceForUserAndFilter indicates an expected call of CountInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) CountInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).CountInstanceForUserAndFilter), ctx, arg)
}

// CreateInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) CreateInstanceForUserAndFilter(ctx context.Context, arg db.CreateInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInstanceForUserAndFilter indicates an expected call of CreateInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) CreateInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).CreateInstanceForUserAndFilter), ctx, arg)
}

// CreateListForUser mocks base method.
func (m *MockQuerier) CreateListForUser(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateListForUser", ctx, userID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateListForUser indicates an expected call of CreateListForUser.
func (mr *MockQuerierMockRecorder) CreateListForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListForUser", reflect.TypeOf((*MockQuerier)(nil).CreateListForUser), ctx, userID)
}

// DeleteInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) DeleteInstanceForUserAndFilter(ctx context.Context, arg db.DeleteInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstanceForUserAndFilter indicates an expected call of DeleteInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) DeleteInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).DeleteInstanceForUserAndFilter), ctx, arg)
}

// GetActiveFiltersForUser mocks base method.
func (m *MockQuerier) GetActiveFiltersForUser(ctx context.Context, userID uuid.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveFiltersForUser", ctx, userID)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveFiltersForUser indicates an expected call of GetActiveFiltersForUser.
func (mr *MockQuerierMockRecorder) GetActiveFiltersForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveFiltersForUser", reflect.TypeOf((*MockQuerier)(nil).GetActiveFiltersForUser), ctx, userID)
}

// GetInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) GetInstanceForUserAndFilter(ctx context.Context, arg db.GetInstanceForUserAndFilterParams) (pgtype.JSONB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(pgtype.JSONB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceForUserAndFilter indicates an expected call of GetInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) GetInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).GetInstanceForUserAndFilter), ctx, arg)
}

// GetInstancesForList mocks base method.
func (m *MockQuerier) GetInstancesForList(ctx context.Context, filterListID int32) ([]db.GetInstancesForListRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstancesForList", ctx, filterListID)
	ret0, _ := ret[0].([]db.GetInstancesForListRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstancesForList indicates an expected call of GetInstancesForList.
func (mr *MockQuerierMockRecorder) GetInstancesForList(ctx, filterListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstancesForList", reflect.TypeOf((*MockQuerier)(nil).GetInstancesForList), ctx, filterListID)
}

// GetListForToken mocks base method.
func (m *MockQuerier) GetListForToken(ctx context.Context, token uuid.UUID) (db.GetListForTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListForToken", ctx, token)
	ret0, _ := ret[0].(db.GetListForTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListForToken indicates an expected call of GetListForToken.
func (mr *MockQuerierMockRecorder) GetListForToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListForToken", reflect.TypeOf((*MockQuerier)(nil).GetListForToken), ctx, token)
}

// GetListForUser mocks base method.
func (m *MockQuerier) GetListForUser(ctx context.Context, userID uuid.UUID) (db.GetListForUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListForUser", ctx, userID)
	ret0, _ := ret[0].(db.GetListForUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListForUser indicates an expected call of GetListForUser.
func (mr *MockQuerierMockRecorder) GetListForUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListForUser", reflect.TypeOf((*MockQuerier)(nil).GetListForUser), ctx, userID)
}

// GetStats mocks base method.
func (m *MockQuerier) GetStats(ctx context.Context) (db.GetStatsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", ctx)
	ret0, _ := ret[0].(db.GetStatsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockQuerierMockRecorder) GetStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockQuerier)(nil).GetStats), ctx)
}

// MarkListDownloaded mocks base method.
func (m *MockQuerier) MarkListDownloaded(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkListDownloaded", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkListDownloaded indicates an expected call of MarkListDownloaded.
func (mr *MockQuerierMockRecorder) MarkListDownloaded(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkListDownloaded", reflect.TypeOf((*MockQuerier)(nil).MarkListDownloaded), ctx, id)
}

// UpdateInstanceForUserAndFilter mocks base method.
func (m *MockQuerier) UpdateInstanceForUserAndFilter(ctx context.Context, arg db.UpdateInstanceForUserAndFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstanceForUserAndFilter", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstanceForUserAndFilter indicates an expected call of UpdateInstanceForUserAndFilter.
func (mr *MockQuerierMockRecorder) UpdateInstanceForUserAndFilter(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstanceForUserAndFilter", reflect.TypeOf((*MockQuerier)(nil).UpdateInstanceForUserAndFilter), ctx, arg)
}