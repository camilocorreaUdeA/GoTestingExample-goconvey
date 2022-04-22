// Code generated by MockGen. DO NOT EDIT.
// Source: ./main.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomClient is a mock of CustomClient interface.
type MockCustomClient struct {
	ctrl     *gomock.Controller
	recorder *MockCustomClientMockRecorder
}

// MockCustomClientMockRecorder is the mock recorder for MockCustomClient.
type MockCustomClientMockRecorder struct {
	mock *MockCustomClient
}

// NewMockCustomClient creates a new mock instance.
func NewMockCustomClient(ctrl *gomock.Controller) *MockCustomClient {
	mock := &MockCustomClient{ctrl: ctrl}
	mock.recorder = &MockCustomClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomClient) EXPECT() *MockCustomClientMockRecorder {
	return m.recorder
}

// GetBoolResponse mocks base method.
func (m *MockCustomClient) GetBoolResponse(option int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoolResponse", option)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBoolResponse indicates an expected call of GetBoolResponse.
func (mr *MockCustomClientMockRecorder) GetBoolResponse(option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoolResponse", reflect.TypeOf((*MockCustomClient)(nil).GetBoolResponse), option)
}

// GetStringResponse mocks base method.
func (m *MockCustomClient) GetStringResponse(option bool) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringResponse", option)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStringResponse indicates an expected call of GetStringResponse.
func (mr *MockCustomClientMockRecorder) GetStringResponse(option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringResponse", reflect.TypeOf((*MockCustomClient)(nil).GetStringResponse), option)
}

// GetntegerResponse mocks base method.
func (m *MockCustomClient) GetntegerResponse(option string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetntegerResponse", option)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetntegerResponse indicates an expected call of GetntegerResponse.
func (mr *MockCustomClientMockRecorder) GetntegerResponse(option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetntegerResponse", reflect.TypeOf((*MockCustomClient)(nil).GetntegerResponse), option)
}