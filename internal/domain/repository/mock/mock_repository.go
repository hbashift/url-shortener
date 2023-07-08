// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/hbashift/url-shortener/internal/domain/repository/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetUrl mocks base method.
func (m *MockRepository) GetUrl(url *model.Url, byLongUrl bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", url, byLongUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockRepositoryMockRecorder) GetUrl(url, byLongUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockRepository)(nil).GetUrl), url, byLongUrl)
}

// PostUrl mocks base method.
func (m *MockRepository) PostUrl(url *model.Url) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUrl", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostUrl indicates an expected call of PostUrl.
func (mr *MockRepositoryMockRecorder) PostUrl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUrl", reflect.TypeOf((*MockRepository)(nil).PostUrl), url)
}
