// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import domain "signIn/gateway/domain"
import mock "github.com/stretchr/testify/mock"

// SignInService is an autogenerated mock type for the SignInService type
type SignInService struct {
	mock.Mock
}

// SignInWithFb provides a mock function with given fields: _a0, _a1
func (_m *SignInService) SignInWithFb(_a0 context.Context, _a1 domain.AccessData) (domain.SignInData, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.SignInData
	if rf, ok := ret.Get(0).(func(context.Context, domain.AccessData) domain.SignInData); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.SignInData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.AccessData) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignInWithGoogle provides a mock function with given fields: _a0, _a1
func (_m *SignInService) SignInWithGoogle(_a0 context.Context, _a1 domain.AccessData) (domain.SignInData, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.SignInData
	if rf, ok := ret.Get(0).(func(context.Context, domain.AccessData) domain.SignInData); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.SignInData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.AccessData) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignInWithLine provides a mock function with given fields: _a0, _a1
func (_m *SignInService) SignInWithLine(_a0 context.Context, _a1 domain.AccessData) (domain.SignInData, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.SignInData
	if rf, ok := ret.Get(0).(func(context.Context, domain.AccessData) domain.SignInData); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.SignInData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.AccessData) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
