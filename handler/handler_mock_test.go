// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	models "github.com/serhiihuberniuk/unstable-api/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockfetcher is a mock of fetcher interface.
type Mockfetcher struct {
	ctrl     *gomock.Controller
	recorder *MockfetcherMockRecorder
}

// MockfetcherMockRecorder is the mock recorder for Mockfetcher.
type MockfetcherMockRecorder struct {
	mock *Mockfetcher
}

// NewMockfetcher creates a new mock instance.
func NewMockfetcher(ctrl *gomock.Controller) *Mockfetcher {
	mock := &Mockfetcher{ctrl: ctrl}
	mock.recorder = &MockfetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockfetcher) EXPECT() *MockfetcherMockRecorder {
	return m.recorder
}

// Leagues mocks base method.
func (m *Mockfetcher) Leagues(ctx context.Context) ([]models.Leagues, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leagues", ctx)
	ret0, _ := ret[0].([]models.Leagues)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leagues indicates an expected call of Leagues.
func (mr *MockfetcherMockRecorder) Leagues(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leagues", reflect.TypeOf((*Mockfetcher)(nil).Leagues), ctx)
}

// Teams mocks base method.
func (m *Mockfetcher) Teams(ctx context.Context) ([]models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teams", ctx)
	ret0, _ := ret[0].([]models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teams indicates an expected call of Teams.
func (mr *MockfetcherMockRecorder) Teams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teams", reflect.TypeOf((*Mockfetcher)(nil).Teams), ctx)
}
