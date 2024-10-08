// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/evcc-io/evcc/server/db/settings (interfaces: API)
//
// Generated by this command:
//
//	mockgen -package settings -destination mock.go -mock_names API=MockAPI github.com/evcc-io/evcc/server/db/settings API
//

// Package settings is a generated GoMock package.
package settings

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// SetString mocks base method.
func (m *MockAPI) SetString(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetString", arg0, arg1)
}

// SetString indicates an expected call of SetString.
func (mr *MockAPIMockRecorder) SetString(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetString", reflect.TypeOf((*MockAPI)(nil).SetString), arg0, arg1)
}

// String mocks base method.
func (m *MockAPI) String(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockAPIMockRecorder) String(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockAPI)(nil).String), arg0)
}
