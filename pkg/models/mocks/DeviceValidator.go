// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	models "github.com/edgexfoundry/go-mod-core-contracts/v3/models"
	mock "github.com/stretchr/testify/mock"
)

// DeviceValidator is an autogenerated mock type for the DeviceValidator type
type DeviceValidator struct {
	mock.Mock
}

// ValidateDevice provides a mock function with given fields: device
func (_m *DeviceValidator) ValidateDevice(device models.Device) error {
	ret := _m.Called(device)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDeviceValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeviceValidator creates a new instance of DeviceValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeviceValidator(t mockConstructorTestingTNewDeviceValidator) *DeviceValidator {
	mock := &DeviceValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}