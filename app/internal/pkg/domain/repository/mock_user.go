// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Change mocks base method.
func (m *MockIUserRepository) Change(ctx context.Context, u entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Change", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Change indicates an expected call of Change.
func (mr *MockIUserRepositoryMockRecorder) Change(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Change", reflect.TypeOf((*MockIUserRepository)(nil).Change), ctx, u)
}

// IsExist mocks base method.
func (m *MockIUserRepository) IsExist(ctx context.Context, userId entity.Id) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExist indicates an expected call of IsExist.
func (mr *MockIUserRepositoryMockRecorder) IsExist(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockIUserRepository)(nil).IsExist), ctx, userId)
}

// List mocks base method.
func (m *MockIUserRepository) List(ctx context.Context, limit entity.Limit, page entity.Page) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, page)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIUserRepositoryMockRecorder) List(ctx, limit, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIUserRepository)(nil).List), ctx, limit, page)
}

// Register mocks base method.
func (m *MockIUserRepository) Register(ctx context.Context, u entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockIUserRepositoryMockRecorder) Register(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIUserRepository)(nil).Register), ctx, u)
}

// SearchByNamePrefix mocks base method.
func (m *MockIUserRepository) SearchByNamePrefix(ctx context.Context, limit entity.Limit, page entity.Page, un entity.UserName) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByNamePrefix", ctx, limit, page, un)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByNamePrefix indicates an expected call of SearchByNamePrefix.
func (mr *MockIUserRepositoryMockRecorder) SearchByNamePrefix(ctx, limit, page, un interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByNamePrefix", reflect.TypeOf((*MockIUserRepository)(nil).SearchByNamePrefix), ctx, limit, page, un)
}

// Select mocks base method.
func (m *MockIUserRepository) Select(ctx context.Context, limit entity.Limit, page entity.Page, ids entity.Ids) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx, limit, page, ids)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockIUserRepositoryMockRecorder) Select(ctx, limit, page, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIUserRepository)(nil).Select), ctx, limit, page, ids)
}

// Self mocks base method.
func (m *MockIUserRepository) Self(ctx context.Context) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Self", ctx)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Self indicates an expected call of Self.
func (mr *MockIUserRepositoryMockRecorder) Self(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockIUserRepository)(nil).Self), ctx)
}
