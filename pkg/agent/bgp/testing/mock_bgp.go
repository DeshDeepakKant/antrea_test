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
// Source: antrea.io/antrea/pkg/agent/bgp (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/agent/bgp/testing/mock_bgp.go -package testing antrea.io/antrea/pkg/agent/bgp Interface
//

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	bgp "antrea.io/antrea/pkg/agent/bgp"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AddPeer mocks base method.
func (m *MockInterface) AddPeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPeer", ctx, peerConf)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeer indicates an expected call of AddPeer.
func (mr *MockInterfaceMockRecorder) AddPeer(ctx, peerConf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockInterface)(nil).AddPeer), ctx, peerConf)
}

// AdvertiseRoutes mocks base method.
func (m *MockInterface) AdvertiseRoutes(ctx context.Context, routes []bgp.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdvertiseRoutes", ctx, routes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdvertiseRoutes indicates an expected call of AdvertiseRoutes.
func (mr *MockInterfaceMockRecorder) AdvertiseRoutes(ctx, routes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdvertiseRoutes", reflect.TypeOf((*MockInterface)(nil).AdvertiseRoutes), ctx, routes)
}

// GetPeers mocks base method.
func (m *MockInterface) GetPeers(ctx context.Context) ([]bgp.PeerStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeers", ctx)
	ret0, _ := ret[0].([]bgp.PeerStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeers indicates an expected call of GetPeers.
func (mr *MockInterfaceMockRecorder) GetPeers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeers", reflect.TypeOf((*MockInterface)(nil).GetPeers), ctx)
}

// GetRoutes mocks base method.
func (m *MockInterface) GetRoutes(ctx context.Context, routeType bgp.RouteType, peerAddress string) ([]bgp.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutes", ctx, routeType, peerAddress)
	ret0, _ := ret[0].([]bgp.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutes indicates an expected call of GetRoutes.
func (mr *MockInterfaceMockRecorder) GetRoutes(ctx, routeType, peerAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutes", reflect.TypeOf((*MockInterface)(nil).GetRoutes), ctx, routeType, peerAddress)
}

// RemovePeer mocks base method.
func (m *MockInterface) RemovePeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePeer", ctx, peerConf)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePeer indicates an expected call of RemovePeer.
func (mr *MockInterfaceMockRecorder) RemovePeer(ctx, peerConf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePeer", reflect.TypeOf((*MockInterface)(nil).RemovePeer), ctx, peerConf)
}

// Start mocks base method.
func (m *MockInterface) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockInterfaceMockRecorder) Start(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInterface)(nil).Start), ctx)
}

// Stop mocks base method.
func (m *MockInterface) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockInterfaceMockRecorder) Stop(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInterface)(nil).Stop), ctx)
}

// UpdatePeer mocks base method.
func (m *MockInterface) UpdatePeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeer", ctx, peerConf)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePeer indicates an expected call of UpdatePeer.
func (mr *MockInterfaceMockRecorder) UpdatePeer(ctx, peerConf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeer", reflect.TypeOf((*MockInterface)(nil).UpdatePeer), ctx, peerConf)
}

// WithdrawRoutes mocks base method.
func (m *MockInterface) WithdrawRoutes(ctx context.Context, routes []bgp.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawRoutes", ctx, routes)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawRoutes indicates an expected call of WithdrawRoutes.
func (mr *MockInterfaceMockRecorder) WithdrawRoutes(ctx, routes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawRoutes", reflect.TypeOf((*MockInterface)(nil).WithdrawRoutes), ctx, routes)
}
