// Code generated by MockGen. DO NOT EDIT.
// Source: entity.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	city "github.com/buni/aliens-game/internal/pkg/board/city"
	player "github.com/buni/aliens-game/internal/pkg/player"
	gomock "github.com/golang/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// DeleteCity mocks base method.
func (m *MockState) DeleteCity(cityName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCity", cityName)
}

// DeleteCity indicates an expected call of DeleteCity.
func (mr *MockStateMockRecorder) DeleteCity(cityName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCity", reflect.TypeOf((*MockState)(nil).DeleteCity), cityName)
}

// DeleteCityAndLinks mocks base method.
func (m *MockState) DeleteCityAndLinks(cityName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCityAndLinks", cityName)
}

// DeleteCityAndLinks indicates an expected call of DeleteCityAndLinks.
func (mr *MockStateMockRecorder) DeleteCityAndLinks(cityName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCityAndLinks", reflect.TypeOf((*MockState)(nil).DeleteCityAndLinks), cityName)
}

// GetCities mocks base method.
func (m *MockState) GetCities() []*city.City {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCities")
	ret0, _ := ret[0].([]*city.City)
	return ret0
}

// GetCities indicates an expected call of GetCities.
func (mr *MockStateMockRecorder) GetCities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCities", reflect.TypeOf((*MockState)(nil).GetCities))
}

// GetCity mocks base method.
func (m *MockState) GetCity(cityName string) (*city.City, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCity", cityName)
	ret0, _ := ret[0].(*city.City)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCity indicates an expected call of GetCity.
func (mr *MockStateMockRecorder) GetCity(cityName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCity", reflect.TypeOf((*MockState)(nil).GetCity), cityName)
}

// GetNextDirection mocks base method.
func (m *MockState) GetNextDirection(currentCity string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextDirection", currentCity)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextDirection indicates an expected call of GetNextDirection.
func (mr *MockStateMockRecorder) GetNextDirection(currentCity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextDirection", reflect.TypeOf((*MockState)(nil).GetNextDirection), currentCity)
}

// MoveVisitor mocks base method.
func (m *MockState) MoveVisitor(cityName string, visitor player.Player) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveVisitor", cityName, visitor)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MoveVisitor indicates an expected call of MoveVisitor.
func (mr *MockStateMockRecorder) MoveVisitor(cityName, visitor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveVisitor", reflect.TypeOf((*MockState)(nil).MoveVisitor), cityName, visitor)
}
