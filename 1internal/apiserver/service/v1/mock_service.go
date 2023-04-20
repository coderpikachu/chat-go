// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by MockGen. DO NOT EDIT.
// Source: chat-go/1internal/apiserver/service/v1 (interfaces: Service,UserSrv,SecretSrv,PolicySrv)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "chat-go/4api/apiserver/v1"
	v10 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Users mocks base method.
func (m *MockService) Users() UserSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(UserSrv)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockServiceMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockService)(nil).Users))
}

// MockUserSrv is a mock of UserSrv interface.
type MockUserSrv struct {
	ctrl     *gomock.Controller
	recorder *MockUserSrvMockRecorder
}

// MockUserSrvMockRecorder is the mock recorder for MockUserSrv.
type MockUserSrvMockRecorder struct {
	mock *MockUserSrv
}

// NewMockUserSrv creates a new mock instance.
func NewMockUserSrv(ctrl *gomock.Controller) *MockUserSrv {
	mock := &MockUserSrv{ctrl: ctrl}
	mock.recorder = &MockUserSrvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSrv) EXPECT() *MockUserSrvMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserSrv) ChangePassword(arg0 context.Context, arg1 *v1.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserSrvMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserSrv)(nil).ChangePassword), arg0, arg1)
}

// Create mocks base method.
func (m *MockUserSrv) Create(arg0 context.Context, arg1 *v1.User, arg2 v10.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserSrvMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserSrv)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockUserSrv) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserSrvMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserSrv)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockUserSrv) DeleteCollection(arg0 context.Context, arg1 []string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockUserSrvMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockUserSrv)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockUserSrv) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserSrvMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserSrv)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockUserSrv) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserSrvMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserSrv)(nil).List), arg0, arg1)
}

// ListWithBadPerformance mocks base method.
func (m *MockUserSrv) ListWithBadPerformance(arg0 context.Context, arg1 v10.ListOptions) (*v1.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWithBadPerformance", arg0, arg1)
	ret0, _ := ret[0].(*v1.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWithBadPerformance indicates an expected call of ListWithBadPerformance.
func (mr *MockUserSrvMockRecorder) ListWithBadPerformance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWithBadPerformance", reflect.TypeOf((*MockUserSrv)(nil).ListWithBadPerformance), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserSrv) Update(arg0 context.Context, arg1 *v1.User, arg2 v10.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserSrvMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserSrv)(nil).Update), arg0, arg1, arg2)
}
