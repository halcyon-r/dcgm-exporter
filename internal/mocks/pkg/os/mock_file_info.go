// Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: io/fs (interfaces: FileInfo)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/pkg/os/mock_file_info.go -package=os -copyright_file=../../../hack/header.txt io/fs FileInfo
//

// Package os is a generated GoMock package.
package os

import (
	fs "io/fs"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockFileInfo is a mock of FileInfo interface.
type MockFileInfo struct {
	ctrl     *gomock.Controller
	recorder *MockFileInfoMockRecorder
	isgomock struct{}
}

// MockFileInfoMockRecorder is the mock recorder for MockFileInfo.
type MockFileInfoMockRecorder struct {
	mock *MockFileInfo
}

// NewMockFileInfo creates a new mock instance.
func NewMockFileInfo(ctrl *gomock.Controller) *MockFileInfo {
	mock := &MockFileInfo{ctrl: ctrl}
	mock.recorder = &MockFileInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileInfo) EXPECT() *MockFileInfoMockRecorder {
	return m.recorder
}

// IsDir mocks base method.
func (m *MockFileInfo) IsDir() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir.
func (mr *MockFileInfoMockRecorder) IsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockFileInfo)(nil).IsDir))
}

// ModTime mocks base method.
func (m *MockFileInfo) ModTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// ModTime indicates an expected call of ModTime.
func (mr *MockFileInfoMockRecorder) ModTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModTime", reflect.TypeOf((*MockFileInfo)(nil).ModTime))
}

// Mode mocks base method.
func (m *MockFileInfo) Mode() fs.FileMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(fs.FileMode)
	return ret0
}

// Mode indicates an expected call of Mode.
func (mr *MockFileInfoMockRecorder) Mode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockFileInfo)(nil).Mode))
}

// Name mocks base method.
func (m *MockFileInfo) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFileInfoMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFileInfo)(nil).Name))
}

// Size mocks base method.
func (m *MockFileInfo) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockFileInfoMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockFileInfo)(nil).Size))
}

// Sys mocks base method.
func (m *MockFileInfo) Sys() any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sys")
	ret0, _ := ret[0].(any)
	return ret0
}

// Sys indicates an expected call of Sys.
func (mr *MockFileInfoMockRecorder) Sys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sys", reflect.TypeOf((*MockFileInfo)(nil).Sys))
}
