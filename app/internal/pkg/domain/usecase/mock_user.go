// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	model "github.com/Khmer495/go-templete/internal/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// ChangeSelfProfile mocks base method.
func (m *MockIUserUsecase) ChangeSelfProfile(ctx context.Context, pName *string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSelfProfile", ctx, pName)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeSelfProfile indicates an expected call of ChangeSelfProfile.
func (mr *MockIUserUsecaseMockRecorder) ChangeSelfProfile(ctx, pName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSelfProfile", reflect.TypeOf((*MockIUserUsecase)(nil).ChangeSelfProfile), ctx, pName)
}

// GetList mocks base method.
func (m *MockIUserUsecase) GetList(ctx context.Context, limit, page int, pName *string) (model.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, limit, page, pName)
	ret0, _ := ret[0].(model.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockIUserUsecaseMockRecorder) GetList(ctx, limit, page, pName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockIUserUsecase)(nil).GetList), ctx, limit, page, pName)
}

// GetSelfProfile mocks base method.
func (m *MockIUserUsecase) GetSelfProfile(ctx context.Context) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfProfile", ctx)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelfProfile indicates an expected call of GetSelfProfile.
func (mr *MockIUserUsecaseMockRecorder) GetSelfProfile(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfProfile", reflect.TypeOf((*MockIUserUsecase)(nil).GetSelfProfile), ctx)
}

// Register mocks base method.
func (m *MockIUserUsecase) Register(ctx context.Context, name string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, name)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockIUserUsecaseMockRecorder) Register(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIUserUsecase)(nil).Register), ctx, name)
}
