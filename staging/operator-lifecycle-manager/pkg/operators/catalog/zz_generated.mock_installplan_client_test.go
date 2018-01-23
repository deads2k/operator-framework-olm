// Automatically generated by MockGen. DO NOT EDIT!
// Source: pkg/client/installplan_client.go

package catalog

import (
	v1alpha1 "github.com/coreos-inc/alm/pkg/apis/installplan/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// Mock of InstallPlanInterface interface
type MockInstallPlanInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockInstallPlanInterfaceRecorder
}

// Recorder for MockInstallPlanInterface (not exported)
type _MockInstallPlanInterfaceRecorder struct {
	mock *MockInstallPlanInterface
}

func NewMockInstallPlanInterface(ctrl *gomock.Controller) *MockInstallPlanInterface {
	mock := &MockInstallPlanInterface{ctrl: ctrl}
	mock.recorder = &_MockInstallPlanInterfaceRecorder{mock}
	return mock
}

func (_m *MockInstallPlanInterface) EXPECT() *_MockInstallPlanInterfaceRecorder {
	return _m.recorder
}

func (_m *MockInstallPlanInterface) UpdateInstallPlan(_param0 *v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error) {
	ret := _m.ctrl.Call(_m, "UpdateInstallPlan", _param0)
	ret0, _ := ret[0].(*v1alpha1.InstallPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallPlanInterfaceRecorder) UpdateInstallPlan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateInstallPlan", arg0)
}

func (_m *MockInstallPlanInterface) CreateInstallPlan(_param0 *v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error) {
	ret := _m.ctrl.Call(_m, "CreateInstallPlan", _param0)
	ret0, _ := ret[0].(*v1alpha1.InstallPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstallPlanInterfaceRecorder) CreateInstallPlan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateInstallPlan", arg0)
}
