// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "signIn/line/domain"
import mock "github.com/stretchr/testify/mock"

// SignInService is an autogenerated mock type for the SignInService type
type SignInService struct {
	mock.Mock
}

// SignIn provides a mock function with given fields: _a0
func (_m *SignInService) SignIn(_a0 string) (domain.SignInResponse, error) {
	ret := _m.Called(_a0)

	var r0 domain.SignInResponse
	if rf, ok := ret.Get(0).(func(string) domain.SignInResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.SignInResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
