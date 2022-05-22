// Code generated by MockGen. DO NOT EDIT.
// Source: entity.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPlayer is a mock of Player interface.
type MockPlayer struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerMockRecorder
}

// MockPlayerMockRecorder is the mock recorder for MockPlayer.
type MockPlayerMockRecorder struct {
	mock *MockPlayer
}

// NewMockPlayer creates a new mock instance.
func NewMockPlayer(ctrl *gomock.Controller) *MockPlayer {
	mock := &MockPlayer{ctrl: ctrl}
	mock.recorder = &MockPlayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayer) EXPECT() *MockPlayerMockRecorder {
	return m.recorder
}

// City mocks base method.
func (m *MockPlayer) City() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "City")
	ret0, _ := ret[0].(string)
	return ret0
}

// City indicates an expected call of City.
func (mr *MockPlayerMockRecorder) City() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "City", reflect.TypeOf((*MockPlayer)(nil).City))
}

// Destroy mocks base method.
func (m *MockPlayer) Destroy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy.
func (mr *MockPlayerMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockPlayer)(nil).Destroy))
}

// IsDestroyed mocks base method.
func (m *MockPlayer) IsDestroyed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDestroyed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDestroyed indicates an expected call of IsDestroyed.
func (mr *MockPlayerMockRecorder) IsDestroyed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDestroyed", reflect.TypeOf((*MockPlayer)(nil).IsDestroyed))
}

// Move mocks base method.
func (m *MockPlayer) Move(city string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", city)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockPlayerMockRecorder) Move(city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockPlayer)(nil).Move), city)
}

// Name mocks base method.
func (m *MockPlayer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPlayerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPlayer)(nil).Name))
}

// SetCity mocks base method.
func (m *MockPlayer) SetCity(city string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCity", city)
}

// SetCity indicates an expected call of SetCity.
func (mr *MockPlayerMockRecorder) SetCity(city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCity", reflect.TypeOf((*MockPlayer)(nil).SetCity), city)
}