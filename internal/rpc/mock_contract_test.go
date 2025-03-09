// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package rpc is a generated GoMock package.
package rpc

import (
	context "context"
	reflect "reflect"

	model "github.com/The-Fox-Hunt/auth/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// GetPassword mocks base method.
func (m *MockRepo) GetPassword(ctx context.Context, username string) (model.UserPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", ctx, username)
	ret0, _ := ret[0].(model.UserPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockRepoMockRecorder) GetPassword(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockRepo)(nil).GetPassword), ctx, username)
}

// Insert mocks base method.
func (m *MockRepo) Insert(ctx context.Context, username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepoMockRecorder) Insert(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepo)(nil).Insert), ctx, username, password)
}

// UpdatePassword mocks base method.
func (m *MockRepo) UpdatePassword(ctx context.Context, username string, newPassword model.UserPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, username, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockRepoMockRecorder) UpdatePassword(ctx, username, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockRepo)(nil).UpdatePassword), ctx, username, newPassword)
}
