// Code generated by MockGen. DO NOT EDIT.
// Source: plugin.go

// Package mock_plugin is a generated GoMock package.
package mock_plugin

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
)

// MockVendorPlugin is a mock of VendorPlugin interface.
type MockVendorPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockVendorPluginMockRecorder
}

// MockVendorPluginMockRecorder is the mock recorder for MockVendorPlugin.
type MockVendorPluginMockRecorder struct {
	mock *MockVendorPlugin
}

// NewMockVendorPlugin creates a new mock instance.
func NewMockVendorPlugin(ctrl *gomock.Controller) *MockVendorPlugin {
	mock := &MockVendorPlugin{ctrl: ctrl}
	mock.recorder = &MockVendorPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVendorPlugin) EXPECT() *MockVendorPluginMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockVendorPlugin) Apply() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply")
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockVendorPluginMockRecorder) Apply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockVendorPlugin)(nil).Apply))
}

// CheckStatusChanges mocks base method.
func (m *MockVendorPlugin) CheckStatusChanges(arg0 *v1.SriovNetworkNodeState) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStatusChanges", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckStatusChanges indicates an expected call of CheckStatusChanges.
func (mr *MockVendorPluginMockRecorder) CheckStatusChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStatusChanges", reflect.TypeOf((*MockVendorPlugin)(nil).CheckStatusChanges), arg0)
}

// Name mocks base method.
func (m *MockVendorPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockVendorPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockVendorPlugin)(nil).Name))
}

// OnNodeStateChange mocks base method.
func (m *MockVendorPlugin) OnNodeStateChange(arg0 *v1.SriovNetworkNodeState) (bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnNodeStateChange", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnNodeStateChange indicates an expected call of OnNodeStateChange.
func (mr *MockVendorPluginMockRecorder) OnNodeStateChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNodeStateChange", reflect.TypeOf((*MockVendorPlugin)(nil).OnNodeStateChange), arg0)
}

// Spec mocks base method.
func (m *MockVendorPlugin) Spec() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spec")
	ret0, _ := ret[0].(string)
	return ret0
}

// Spec indicates an expected call of Spec.
func (mr *MockVendorPluginMockRecorder) Spec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spec", reflect.TypeOf((*MockVendorPlugin)(nil).Spec))
}
