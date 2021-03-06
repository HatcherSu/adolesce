// Code generated by MockGen. DO NOT EDIT.
// Source: callback_log.go

// Package biz is a generated GoMock package.
package biz

import (
	biz "adolesce/internal/biz"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCallbackLogRepo is a mock of CallbackLogRepo interface.
type MockCallbackLogRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCallbackLogRepoMockRecorder
}

// MockCallbackLogRepoMockRecorder is the mock recorder for MockCallbackLogRepo.
type MockCallbackLogRepoMockRecorder struct {
	mock *MockCallbackLogRepo
}

// NewMockCallbackLogRepo creates a new mock instance.
func NewMockCallbackLogRepo(ctrl *gomock.Controller) *MockCallbackLogRepo {
	mock := &MockCallbackLogRepo{ctrl: ctrl}
	mock.recorder = &MockCallbackLogRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCallbackLogRepo) EXPECT() *MockCallbackLogRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCallbackLogRepo) Create(arg0 *biz.CallbackLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCallbackLogRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCallbackLogRepo)(nil).Create), arg0)
}

// QueryList mocks base method.
func (m *MockCallbackLogRepo) QueryList(arg0 *biz.CallbackLogFilter) ([]*biz.CallbackLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryList", arg0)
	ret0, _ := ret[0].([]*biz.CallbackLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryList indicates an expected call of QueryList.
func (mr *MockCallbackLogRepoMockRecorder) QueryList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryList", reflect.TypeOf((*MockCallbackLogRepo)(nil).QueryList), arg0)
}
