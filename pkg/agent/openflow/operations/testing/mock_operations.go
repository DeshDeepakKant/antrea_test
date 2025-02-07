// Copyright 2024 Antrea Authors
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
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/openflow/operations (interfaces: OFEntryOperations)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/agent/openflow/operations/testing/mock_operations.go -package testing antrea.io/antrea/pkg/agent/openflow/operations OFEntryOperations
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	openflow "antrea.io/antrea/pkg/ovs/openflow"
	openflow15 "antrea.io/libOpenflow/openflow15"
	gomock "go.uber.org/mock/gomock"
)

// MockOFEntryOperations is a mock of OFEntryOperations interface.
type MockOFEntryOperations struct {
	ctrl     *gomock.Controller
	recorder *MockOFEntryOperationsMockRecorder
	isgomock struct{}
}

// MockOFEntryOperationsMockRecorder is the mock recorder for MockOFEntryOperations.
type MockOFEntryOperationsMockRecorder struct {
	mock *MockOFEntryOperations
}

// NewMockOFEntryOperations creates a new mock instance.
func NewMockOFEntryOperations(ctrl *gomock.Controller) *MockOFEntryOperations {
	mock := &MockOFEntryOperations{ctrl: ctrl}
	mock.recorder = &MockOFEntryOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOFEntryOperations) EXPECT() *MockOFEntryOperationsMockRecorder {
	return m.recorder
}

// AddAll mocks base method.
func (m *MockOFEntryOperations) AddAll(flows []*openflow15.FlowMod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAll", flows)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAll indicates an expected call of AddAll.
func (mr *MockOFEntryOperationsMockRecorder) AddAll(flows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAll", reflect.TypeOf((*MockOFEntryOperations)(nil).AddAll), flows)
}

// AddOFEntries mocks base method.
func (m *MockOFEntryOperations) AddOFEntries(ofEntries []openflow.OFEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOFEntries", ofEntries)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOFEntries indicates an expected call of AddOFEntries.
func (mr *MockOFEntryOperationsMockRecorder) AddOFEntries(ofEntries any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOFEntries", reflect.TypeOf((*MockOFEntryOperations)(nil).AddOFEntries), ofEntries)
}

// BundleOps mocks base method.
func (m *MockOFEntryOperations) BundleOps(adds, mods, dels []*openflow15.FlowMod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BundleOps", adds, mods, dels)
	ret0, _ := ret[0].(error)
	return ret0
}

// BundleOps indicates an expected call of BundleOps.
func (mr *MockOFEntryOperationsMockRecorder) BundleOps(adds, mods, dels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BundleOps", reflect.TypeOf((*MockOFEntryOperations)(nil).BundleOps), adds, mods, dels)
}

// DeleteAll mocks base method.
func (m *MockOFEntryOperations) DeleteAll(flows []*openflow15.FlowMod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", flows)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOFEntryOperationsMockRecorder) DeleteAll(flows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOFEntryOperations)(nil).DeleteAll), flows)
}

// DeleteOFEntries mocks base method.
func (m *MockOFEntryOperations) DeleteOFEntries(ofEntries []openflow.OFEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOFEntries", ofEntries)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOFEntries indicates an expected call of DeleteOFEntries.
func (mr *MockOFEntryOperationsMockRecorder) DeleteOFEntries(ofEntries any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOFEntries", reflect.TypeOf((*MockOFEntryOperations)(nil).DeleteOFEntries), ofEntries)
}

// ModifyAll mocks base method.
func (m *MockOFEntryOperations) ModifyAll(flows []*openflow15.FlowMod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyAll", flows)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyAll indicates an expected call of ModifyAll.
func (mr *MockOFEntryOperationsMockRecorder) ModifyAll(flows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyAll", reflect.TypeOf((*MockOFEntryOperations)(nil).ModifyAll), flows)
}

// ModifyOFEntries mocks base method.
func (m *MockOFEntryOperations) ModifyOFEntries(ofEntries []openflow.OFEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyOFEntries", ofEntries)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyOFEntries indicates an expected call of ModifyOFEntries.
func (mr *MockOFEntryOperationsMockRecorder) ModifyOFEntries(ofEntries any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyOFEntries", reflect.TypeOf((*MockOFEntryOperations)(nil).ModifyOFEntries), ofEntries)
}
