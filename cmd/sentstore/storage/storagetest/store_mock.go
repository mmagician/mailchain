// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package storagetest is a generated GoMock package.
package storagetest

import (
	gomock "github.com/golang/mock/gomock"
	mail "github.com/mailchain/mailchain/internal/mail"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockStore) Exists(messageID mail.ID, contentsHash, integrityHash, contents []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", messageID, contentsHash, integrityHash, contents)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockStoreMockRecorder) Exists(messageID, contentsHash, integrityHash, contents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockStore)(nil).Exists), messageID, contentsHash, integrityHash, contents)
}

// Put mocks base method
func (m *MockStore) Put(messageID mail.ID, contentsHash, integrityHash, contents []byte) (string, string, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", messageID, contentsHash, integrityHash, contents)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Put indicates an expected call of Put
func (mr *MockStoreMockRecorder) Put(messageID, contentsHash, integrityHash, contents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStore)(nil).Put), messageID, contentsHash, integrityHash, contents)
}
