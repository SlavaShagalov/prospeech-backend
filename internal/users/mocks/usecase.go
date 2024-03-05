// Code generated by MockGen. DO NOT EDIT.
// Source: internal/users/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/SlavaShagalov/prospeech-backend/internal/models"
	users "github.com/SlavaShagalov/prospeech-backend/internal/users"
	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUsecase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsecase)(nil).Delete), id)
}

// FullUpdate mocks base method.
func (m *MockUsecase) FullUpdate(params *users.FullUpdateParams) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullUpdate", params)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FullUpdate indicates an expected call of FullUpdate.
func (mr *MockUsecaseMockRecorder) FullUpdate(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullUpdate", reflect.TypeOf((*MockUsecase)(nil).FullUpdate), params)
}

// Get mocks base method.
func (m *MockUsecase) Get(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUsecaseMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsecase)(nil).Get), id)
}

// GetByUsername mocks base method.
func (m *MockUsecase) GetByUsername(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUsecaseMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUsecase)(nil).GetByUsername), username)
}

// List mocks base method.
func (m *MockUsecase) List() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsecaseMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsecase)(nil).List))
}

// PartialUpdate mocks base method.
func (m *MockUsecase) PartialUpdate(params *users.PartialUpdateParams) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PartialUpdate", params)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PartialUpdate indicates an expected call of PartialUpdate.
func (mr *MockUsecaseMockRecorder) PartialUpdate(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PartialUpdate", reflect.TypeOf((*MockUsecase)(nil).PartialUpdate), params)
}

// UpdateAvatar mocks base method.
func (m *MockUsecase) UpdateAvatar(id int, imgData []byte, filename string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", id, imgData, filename)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockUsecaseMockRecorder) UpdateAvatar(id, imgData, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUsecase)(nil).UpdateAvatar), id, imgData, filename)
}
