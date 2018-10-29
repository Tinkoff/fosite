// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/pkce (interfaces: PKCERequestStorage)

package internal

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of PKCERequestStorage interface
type MockPKCERequestStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockPKCERequestStorageRecorder
}

// Recorder for MockPKCERequestStorage (not exported)
type _MockPKCERequestStorageRecorder struct {
	mock *MockPKCERequestStorage
}

func NewMockPKCERequestStorage(ctrl *gomock.Controller) *MockPKCERequestStorage {
	mock := &MockPKCERequestStorage{ctrl: ctrl}
	mock.recorder = &_MockPKCERequestStorageRecorder{mock}
	return mock
}

func (_m *MockPKCERequestStorage) EXPECT() *_MockPKCERequestStorageRecorder {
	return _m.recorder
}

func (_m *MockPKCERequestStorage) CreatePKCERequestSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreatePKCERequestSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPKCERequestStorageRecorder) CreatePKCERequestSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreatePKCERequestSession", arg0, arg1, arg2)
}

func (_m *MockPKCERequestStorage) DeletePKCERequestSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeletePKCERequestSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPKCERequestStorageRecorder) DeletePKCERequestSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeletePKCERequestSession", arg0, arg1)
}

func (_m *MockPKCERequestStorage) GetPKCERequestSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetPKCERequestSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPKCERequestStorageRecorder) GetPKCERequestSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPKCERequestSession", arg0, arg1, arg2)
}
