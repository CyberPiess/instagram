// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_post is a generated GoMock package.
package mock_post

import (
	reflect "reflect"

	database "github.com/CyberPiess/instagram/internal/infrastructure/database/post"
	post "github.com/CyberPiess/instagram/internal/infrastructure/minio/post"
	token "github.com/CyberPiess/instagram/internal/infrastructure/token"
	gomock "github.com/golang/mock/gomock"
)

// MockpostStorage is a mock of postStorage interface.
type MockpostStorage struct {
	ctrl     *gomock.Controller
	recorder *MockpostStorageMockRecorder
}

// MockpostStorageMockRecorder is the mock recorder for MockpostStorage.
type MockpostStorageMockRecorder struct {
	mock *MockpostStorage
}

// NewMockpostStorage creates a new mock instance.
func NewMockpostStorage(ctrl *gomock.Controller) *MockpostStorage {
	mock := &MockpostStorage{ctrl: ctrl}
	mock.recorder = &MockpostStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpostStorage) EXPECT() *MockpostStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockpostStorage) Create(CreatePost database.CreatePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", CreatePost)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockpostStorageMockRecorder) Create(CreatePost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockpostStorage)(nil).Create), CreatePost)
}

// MocktokenInteraction is a mock of tokenInteraction interface.
type MocktokenInteraction struct {
	ctrl     *gomock.Controller
	recorder *MocktokenInteractionMockRecorder
}

// MocktokenInteractionMockRecorder is the mock recorder for MocktokenInteraction.
type MocktokenInteractionMockRecorder struct {
	mock *MocktokenInteraction
}

// NewMocktokenInteraction creates a new mock instance.
func NewMocktokenInteraction(ctrl *gomock.Controller) *MocktokenInteraction {
	mock := &MocktokenInteraction{ctrl: ctrl}
	mock.recorder = &MocktokenInteractionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktokenInteraction) EXPECT() *MocktokenInteractionMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MocktokenInteraction) CreateToken(userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MocktokenInteractionMockRecorder) CreateToken(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MocktokenInteraction)(nil).CreateToken), userId)
}

// VerifyToken mocks base method.
func (m *MocktokenInteraction) VerifyToken(tokenString string) (*token.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", tokenString)
	ret0, _ := ret[0].(*token.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MocktokenInteractionMockRecorder) VerifyToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MocktokenInteraction)(nil).VerifyToken), tokenString)
}

// MockminioStorage is a mock of minioStorage interface.
type MockminioStorage struct {
	ctrl     *gomock.Controller
	recorder *MockminioStorageMockRecorder
}

// MockminioStorageMockRecorder is the mock recorder for MockminioStorage.
type MockminioStorageMockRecorder struct {
	mock *MockminioStorage
}

// NewMockminioStorage creates a new mock instance.
func NewMockminioStorage(ctrl *gomock.Controller) *MockminioStorage {
	mock := &MockminioStorage{ctrl: ctrl}
	mock.recorder = &MockminioStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockminioStorage) EXPECT() *MockminioStorageMockRecorder {
	return m.recorder
}

// UploadFile mocks base method.
func (m *MockminioStorage) UploadFile(arg0 post.ImageDTO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UploadFile", arg0)
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockminioStorageMockRecorder) UploadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockminioStorage)(nil).UploadFile), arg0)
}
