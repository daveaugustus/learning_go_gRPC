// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/davetweetlive/learning_go_gRPC/repos (interfaces: GlobalRepository)

// Package mock_repos is a generated GoMock package.
package mock_repos

import (
	reflect "reflect"

	repos "github.com/davetweetlive/learning_go_gRPC/repos"
	gomock "github.com/golang/mock/gomock"
)

// MockGlobalRepository is a mock of GlobalRepository interface.
type MockGlobalRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalRepositoryMockRecorder
}

// MockGlobalRepositoryMockRecorder is the mock recorder for MockGlobalRepository.
type MockGlobalRepositoryMockRecorder struct {
	mock *MockGlobalRepository
}

// NewMockGlobalRepository creates a new mock instance.
func NewMockGlobalRepository(ctrl *gomock.Controller) *MockGlobalRepository {
	mock := &MockGlobalRepository{ctrl: ctrl}
	mock.recorder = &MockGlobalRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlobalRepository) EXPECT() *MockGlobalRepositoryMockRecorder {
	return m.recorder
}

// Users mocks base method.
func (m *MockGlobalRepository) Users() repos.UsersRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(repos.UsersRepo)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockGlobalRepositoryMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockGlobalRepository)(nil).Users))
}
