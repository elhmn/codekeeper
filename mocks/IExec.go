// Code generated by MockGen. DO NOT EDIT.
// Source: internal/exec/exec.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIExec is a mock of IExec interface.
type MockIExec struct {
	ctrl     *gomock.Controller
	recorder *MockIExecMockRecorder
}

// MockIExecMockRecorder is the mock recorder for MockIExec.
type MockIExecMockRecorder struct {
	mock *MockIExec
}

// NewMockIExec creates a new mock instance.
func NewMockIExec(ctrl *gomock.Controller) *MockIExec {
	mock := &MockIExec{ctrl: ctrl}
	mock.recorder = &MockIExecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExec) EXPECT() *MockIExecMockRecorder {
	return m.recorder
}

// CreateFolderIfDoesNotExist mocks base method.
func (m *MockIExec) CreateFolderIfDoesNotExist(dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolderIfDoesNotExist", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFolderIfDoesNotExist indicates an expected call of CreateFolderIfDoesNotExist.
func (mr *MockIExecMockRecorder) CreateFolderIfDoesNotExist(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolderIfDoesNotExist", reflect.TypeOf((*MockIExec)(nil).CreateFolderIfDoesNotExist), dir)
}

// DoGit mocks base method.
func (m *MockIExec) DoGit(dir string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoGit", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoGit indicates an expected call of DoGit.
func (mr *MockIExecMockRecorder) DoGit(dir interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoGit", reflect.TypeOf((*MockIExec)(nil).DoGit), varargs...)
}

// DoGitClone mocks base method.
func (m *MockIExec) DoGitClone(dir string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoGitClone", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoGitClone indicates an expected call of DoGitClone.
func (mr *MockIExecMockRecorder) DoGitClone(dir interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoGitClone", reflect.TypeOf((*MockIExec)(nil).DoGitClone), varargs...)
}

// DoGitPush mocks base method.
func (m *MockIExec) DoGitPush(dir string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoGitPush", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoGitPush indicates an expected call of DoGitPush.
func (mr *MockIExecMockRecorder) DoGitPush(dir interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoGitPush", reflect.TypeOf((*MockIExec)(nil).DoGitPush), varargs...)
}

// OpenEditor mocks base method.
func (m *MockIExec) OpenEditor(editor string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{editor}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenEditor", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenEditor indicates an expected call of OpenEditor.
func (mr *MockIExecMockRecorder) OpenEditor(editor interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{editor}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenEditor", reflect.TypeOf((*MockIExec)(nil).OpenEditor), varargs...)
}

// Run mocks base method.
func (m *MockIExec) Run(dir, command string, args ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{dir, command}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockIExecMockRecorder) Run(dir, command interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dir, command}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockIExec)(nil).Run), varargs...)
}

// RunInteractive mocks base method.
func (m *MockIExec) RunInteractive(command string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{command}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunInteractive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunInteractive indicates an expected call of RunInteractive.
func (mr *MockIExecMockRecorder) RunInteractive(command interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{command}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInteractive", reflect.TypeOf((*MockIExec)(nil).RunInteractive), varargs...)
}
